package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goClass/global"
	"goClass/model"
	"goClass/model/response"
	"goClass/service"
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
