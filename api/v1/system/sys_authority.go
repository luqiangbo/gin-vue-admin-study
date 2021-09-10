package system

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	commonRes "go-class/model/common/response"
	"go-class/model/system/response"
	"go-class/model/system/tables"
	"go-class/utils"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}

// 创建角色

func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var req tables.SysAuthority
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := authorityService.CreateAuthority(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		commonRes.FailWithMessage(err.Error(), c)
	} else {
		commonRes.OkWithDetailed(response.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// 删除角色
// post

func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var req tables.SysAuthority
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.DeleteAuthority(&req); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		commonRes.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}
