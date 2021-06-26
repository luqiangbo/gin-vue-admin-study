package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goClass/global"
	"goClass/model"
	"goClass/model/request"
	"goClass/model/response"
	"goClass/service"
	"goClass/utils"
)

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

func tokenNext(c *gin.Context, user model.SysUser) {
	response.OkWithData(response.LoginResponse{
		User:  user,
		Token: "123",
	}, c)
}
