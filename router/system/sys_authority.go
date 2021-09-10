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
		authorityRouter.POST("create_authority", authorityApi.CreateAuthority)   // 创建角色(权限)
		authorityRouter.DELETE("delete_authority", authorityApi.DeleteAuthority) // 删除角色(权限)
		authorityRouter.POST("get_authority", authorityApi.GetAuthorityList)     // 分页
		authorityRouter.PUT("update_authority", authorityApi.UpdateAuthority)    // 更新角色
	}
}
