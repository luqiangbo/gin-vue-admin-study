package v1

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	"go-class/model"
	"go-class/model/response"
	"go-class/service"
	"go-class/utils"
	"go.uber.org/zap"
)

func CreateAuthority(c *gin.Context) {
	var req model.SysAuthority
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := service.CreateAuthority(req); err != nil {
		global.GVA_LOG.Error("常见失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}
