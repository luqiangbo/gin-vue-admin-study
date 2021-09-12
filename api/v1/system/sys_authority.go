package system

import (
	"github.com/gin-gonic/gin"
	"gin-vue-admin-study/global"
	commonReq "gin-vue-admin-study/model/common/request"
	commonRes "gin-vue-admin-study/model/common/response"
	"gin-vue-admin-study/model/system/response"
	"gin-vue-admin-study/model/system/tables"
	"gin-vue-admin-study/utils"
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

// 分页获取角色列表

func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var req commonReq.PageInfo
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.PageInfoVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := authorityService.GetAuthorityInfoList(req); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		commonRes.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
			List:     list,
			Total:    total,
			Page:     req.Page,
			PageSize: req.PageSize,
		}, "获取成功", c)
	}
}

// 更新角色

func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var req tables.SysAuthority
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, data := authorityService.UpdateAuthority(req); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		commonRes.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		commonRes.OkWithDetailed(response.SysAuthorityResponse{Authority: data}, "更新成功", c)
	}
}

// 设置角色 资源权限

func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var req tables.SysAuthority
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := authorityService.SetDataAuthority(req); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Any("err", err))
		commonRes.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		commonRes.OkWithMessage("设置成功", c)

	}
}
