package router

import (
	"github.com/gin-gonic/gin"
	v1 "goClass/api/v1"
	"goClass/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)             // 用户注册账号
		UserRouter.PUT("change_password", v1.ChangePassword) // 用户修改密码
		UserRouter.POST("get_list", v1.GetUserList)          //  分页获取用户列表
		UserRouter.DELETE("one", v1.DeleteUser)              // 删除用户
		UserRouter.PUT("info", v1.Info)                      // 更新用户信息
	}
}
