package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goClass/global"
	"goClass/model/request"
	"goClass/model/response"
	"goClass/utils"
)

func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("api", l)
	response.Ok(c)
}
