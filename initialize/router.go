package initialize

import (
	"github.com/gin-gonic/gin"
	"goClass/middleware"
	"goClass/router"
)

// 初始化路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// 方便统一添加路由组签证 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(PrivateGroup) // 注册功能api路由
	}
	return Router
}
