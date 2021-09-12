package system

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	commonRes "go-class/model/common/response"
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
