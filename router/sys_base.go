package router

import (
	"github.com/gin-gonic/gin"
	v1 "goClass/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
	}
	return BaseRouter
}
