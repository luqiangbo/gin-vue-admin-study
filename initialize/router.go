package initialize

import (
	"github.com/gin-gonic/gin"
	"go-class/middleware"
	"go-class/router"
)

// 初始化路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 方便统一添加路由组签证 多服务器上线使用
	// 获取路由实例
	systemRouter := router.RouterGroupApp.System

	// 注册基础功能路由 不做鉴权
	PublicGroup := Router.Group("")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 初始化数据库

	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)      // 注册功能api路由
		systemRouter.InitAuthorityRouter(PrivateGroup) // 权限
	}
	return Router
}
