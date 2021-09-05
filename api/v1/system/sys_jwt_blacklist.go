package system

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	"go-class/model/common/response"
	"go-class/model/system"
	"go.uber.org/zap"
)

type JwtApi struct {
}

func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Any("err", err))
		response.FailWithMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}
}
