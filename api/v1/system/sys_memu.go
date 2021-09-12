package system

import (
	"github.com/gin-gonic/gin"
	"gin-vue-admin-study/global"
	commonReq "gin-vue-admin-study/model/common/request"
	commonRes "gin-vue-admin-study/model/common/response"
	modelSystemReq "gin-vue-admin-study/model/system/request"
	modelSystemRes "gin-vue-admin-study/model/system/response"
	"gin-vue-admin-study/model/system/tables"
	"gin-vue-admin-study/utils"
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

// 增 菜单

func (m *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var req tables.SysBaseMenu
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuDbService.AddBaseMenu(req); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))
		commonRes.FailWithMessage("添加失败", c)

	} else {
		commonRes.OkWithMessage("添加成功", c)
	}
}

// 删 菜单

func (m *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var req commonReq.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuDbService.DeleteBaseMenu(req.ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		commonRes.FailWithMessage("删除失败", c)
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

// 改 菜单

func (m *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var req tables.SysBaseMenu
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.MenuVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(req.Meta, utils.MenuMetaVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuDbService.UpdateBaseMenu(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		commonRes.FailWithMessage("更新失败", c)
	} else {
		commonRes.OkWithMessage("更新成功", c)
	}
}

// 通过id 获取菜单详情

func (m *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var req commonReq.GetById
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.IdVerify); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}
	if err, res := menuDbService.GetBaseMenuById(req.ID); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		commonRes.FailWithMessage("获取失败", c)
	} else {
		commonRes.OkWithDetailed(modelSystemRes.SysBaseMenuResponse{Menu: res}, "获取成功", c)
	}
}
