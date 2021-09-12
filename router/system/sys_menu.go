package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-class/api/v1"
	"go-class/middleware"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	var authorityMenuApi = v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("get_menu", authorityMenuApi.GetMenu) // 获取菜单树
	}
	return menuRouter
}
