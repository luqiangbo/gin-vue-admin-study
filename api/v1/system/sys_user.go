package system

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go-class/global"
	"go-class/middleware"
	commonReq "go-class/model/common/request"
	commonRes "go-class/model/common/response"
	modelSystemRequest "go-class/model/system/request"
	"go-class/model/system/response"
	"go-class/model/system/tables"
	"go-class/utils"
	"go.uber.org/zap"
	"time"
)

type BaseApi struct {
}

// 注册

func (b *BaseApi) Register(c *gin.Context) {
	var req modelSystemRequest.Register
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	user := &tables.SysUser{Username: req.Username, Password: req.Password}
	err, userReturn := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败", zap.Any("err", err))
		commonRes.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		commonRes.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// 登录

func (b *BaseApi) Login(c *gin.Context) {
	var req modelSystemRequest.Login
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	u := &tables.SysUser{Username: req.Username, Password: req.Password}
	if err, user := userService.Login(u); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或密码错误!", zap.Any("err", err))
		commonRes.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

// 获取token
func (b *BaseApi) tokenNext(c *gin.Context, user tables.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}

	claims := modelSystemRequest.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		NickName:   user.NickName,
		Username:   user.Username,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "can",                                                 // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("创建token失败!", zap.Any("err", err))
		commonRes.FailWithMessage("创建token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		commonRes.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	// 根据用户名去redis拿取token
	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		//没找到 就存储
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败1!", zap.Any("err", err))
			commonRes.FailWithMessage("设置登录状态失败1", c)
			return
		}
		//存储成功
		commonRes.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功1", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败2!", zap.Any("err", err))
		commonRes.FailWithMessage("设置登录状态失败2", c)
	} else {
		var blackJWT tables.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			commonRes.FailWithMessage("jwt作废失败3", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			commonRes.FailWithMessage("设置登录状态失败3", c)
			return
		}
		commonRes.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功3", c)
	}
}

// 修改密码

func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req modelSystemRequest.ChangePassword
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.ChangePasswordVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	u := &tables.SysUser{Username: req.Username, Password: req.Password}
	if err, _ := userService.ChangePassword(u, req.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		commonRes.FailWithMessage("修改失败 , 原密码与当前账户不符", c)
	} else {
		commonRes.OkWithMessage("修改成功", c)
	}
}

// 分页用户

func (b *BaseApi) GetUserList(c *gin.Context) {
	var req commonReq.PageInfo
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := userService.GetUserInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		commonRes.FailWithMessage("获取失败", c)
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// 更改用户权限
// @Router /user/set_user_authority post

func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var req modelSystemRequest.SetUserAuth
	_ = c.ShouldBindJSON(&req)
	// 入参校验
	if UserVerifyErr := utils.Verify(req, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		commonRes.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	//
	userID := utils.GetUserId(c)
	uuid := utils.GetUserUuid(c)
	if err := userService.SetUserAuthority(userID, uuid, req.AuthorityId); err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}

// SetUserAuthorities
// 设置用户权限
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var req modelSystemRequest.SetUserAuthorities
	_ = c.ShouldBindJSON(&req)
	if err := userService.SetUserAuthorities(req.ID, req.AuthorityIds); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		commonRes.FailWithMessage("修改失败", c)
	} else {
		commonRes.OkWithMessage("修改成功", c)
	}
}

// 删除用户

func (b *BaseApi) DeleteUser(c *gin.Context) {
	var req commonReq.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := utils.GetUserId(c)
	if jwtId == uint(req.ID) {
		commonRes.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := userService.DeleteUser(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		commonRes.FailWithMessage("删除失败", c)
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

func (b *BaseApi) Info(c *gin.Context) {
	var req tables.SysUser
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := userService.SetUserInfo(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		commonRes.FailWithMessage("设置失败!", c)
	} else {
		commonRes.OkWithDetailed(gin.H{"user_info": res}, "设置成功", c)
	}
}
