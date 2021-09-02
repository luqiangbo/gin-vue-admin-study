package v1

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	"go-class/model"
	"go-class/model/response"
	"go-class/service"
	"go.uber.org/zap"
)

func JsonInBlocklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)

	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
