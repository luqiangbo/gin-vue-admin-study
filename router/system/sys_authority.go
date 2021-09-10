package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-class/api/v1"
	"go-class/middleware"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	var authorityApi = v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("authority", authorityApi.CreateAuthority)          // 创建角色
		authorityRouter.DELETE("delete_authority", authorityApi.DeleteAuthority) // 删除角色
	}
}
