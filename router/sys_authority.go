package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-class/api/v1"
	"go-class/middleware"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	{
		AuthorityRouter.POST("create_authority", v1.CreateAuthority)
	}
}
