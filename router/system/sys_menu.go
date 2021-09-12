package system

import (
	"github.com/gin-gonic/gin"
	v1 "gin-vue-admin-study/api/v1"
	"gin-vue-admin-study/middleware"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	var authorityMenuApi = v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("get_menu", authorityMenuApi.GetMenu)                    // 获取菜单树
		menuRouter.POST("add_menu_authority", authorityMenuApi.AddMenuAuthority) // 增加menu和角色关联关系
		menuRouter.POST("get_menu_authority", authorityMenuApi.GetMenuAuthority) // 获取 指定角色的menu

		menuRouter.POST("add_base_menu", authorityMenuApi.AddBaseMenu)           // 添加角色和menu关系
		menuRouter.DELETE("delete_base_menu", authorityMenuApi.DeleteBaseMenu)   // 删除角色和menu关系
		menuRouter.PUT("update_base_menu", authorityMenuApi.UpdateBaseMenu)      // 更新角色和menu关系
		menuRouter.POST("get_base_menu_tree", authorityMenuApi.GetBaseMenuTree)  // 获取全部菜单
		menuRouter.POST("get_base_menu_by_id", authorityMenuApi.GetBaseMenuById) //
	}
	return menuRouter
}
