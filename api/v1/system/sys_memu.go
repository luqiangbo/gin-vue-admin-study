package system

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	commonReq "go-class/model/common/request"
	commonRes "go-class/model/common/response"
	modelSystemReq "go-class/model/system/request"
	modelSystemRes "go-class/model/system/response"
	"go-class/model/system/tables"
	"go-class/utils"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct {
}

// 获取用户动态路由

func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	if err, menus := menuDbService.GetMenuTree(utils.GetUserAuthorityId(c)); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		commonRes.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []tables.SysMenu{}
		}
		commonRes.OkWithDetailed(modelSystemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// 获取用户的全部路由

func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	if err, menus := menuDbService.GetBaseMenuTree(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		commonRes.FailWithMessage("获取失败", c)
	} else {
		commonRes.OkWithDetailed(modelSystemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// 获取指定角色的菜单

func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param commonReq.GetAuthorityId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := menuDbService.GetMenuAuthority(&param); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		commonRes.FailWithDetailed(modelSystemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		commonRes.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// 增加menu和角色关联关系

func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var req modelSystemReq.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.AuthorityIdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuDbService.AddMenuAuthority(req.Menus, req.AuthorityId); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		commonRes.FailWithMessage("添加失败", c)
	} else {
		commonRes.OkWithMessage("添加成功", c)
	}

}
