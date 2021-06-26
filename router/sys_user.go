package router

import (
	"github.com/gin-gonic/gin"
	v1 "goClass/api/v1"
	"goClass/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register) // 用户注册账号
	}
}
