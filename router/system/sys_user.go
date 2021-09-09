package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-class/api/v1"
	"go-class/middleware"
)

type UserRouter struct {
}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.PUT("change_password", baseApi.ChangePassword)           // 用户修改密码
		userRouter.POST("get_list", baseApi.GetUserList)                    // 分页获取用户列表
		userRouter.DELETE("one", baseApi.DeleteUser)                        // 删除用户
		userRouter.PUT("info", baseApi.Info)                                // 更新用户信息
		userRouter.POST("set_user_authority", baseApi.SetUserAuthority)     // 更改用户权限
		userRouter.POST("set_user_authorities", baseApi.SetUserAuthorities) // 设置用户权限

	}
}
