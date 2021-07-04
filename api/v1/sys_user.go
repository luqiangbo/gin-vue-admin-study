package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"goClass/global"
	"goClass/middleware"
	"goClass/model"
	"goClass/model/request"
	"goClass/model/response"
	"goClass/service"
	"goClass/utils"
	"time"
)

// 注册

func Register(c *gin.Context) {
	var req request.Register
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{Username: req.Username, Password: req.Password}
	err, userReturn := service.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

// 登录

func Login(c *gin.Context) {
	var req request.Login
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.LoginVerify); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &model.SysUser{Username: req.Username, Password: req.Password}
	if err, user := service.Login(u); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或密码错误!", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		tokenNext(c, *user)
	}
}

// 获取token
func tokenNext(c *gin.Context, user model.SysUser) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)}

	claims := request.CustomClaims{
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
		global.GVA_LOG.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	// 根据用户名去redis拿取token
	if err, jwtStr := service.GetRedisJWT(user.Username); err == redis.Nil {
		//没找到 就存储
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败1!", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败1", c)
			return
		}
		//存储成功
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功1", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败2!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败2", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败3", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败3", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功3", c)
	}
}

// 修改密码

func ChangePassword(c *gin.Context) {
	var req request.ChangePassword
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.ChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &model.SysUser{Username: req.Username, Password: req.Password}
	if err, _ := service.ChangePassword(u, req.NewPassword); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败 , 原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// 分页用户

func GetUserList(c *gin.Context) {
	var req request.PageInfo
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetUserInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// 用户权限

func SetUserAuthority(c *gin.Context) {

}

// 删除用户

func DeleteUser(c *gin.Context) {
	var req request.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := getUserId(c)
	if jwtId == uint(req.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	if err := service.DeleteUser(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func Info(c *gin.Context) {
	var req model.SysUser
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := service.SetUserInfo(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		response.FailWithMessage("设置失败!", c)
	} else {
		response.OkWithDetailed(gin.H{"user_info": res}, "设置成功", c)
	}
}

//
func getUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是佛使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}
