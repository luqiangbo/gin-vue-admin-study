package router

import (
	"github.com/gin-gonic/gin"
	v1 "goClass/api/v1"
	"goClass/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("register", v1.Register)              // 用户注册账号
		UserRouter.POST("change_password", v1.ChangePassword) // 用户修改密码
		UserRouter.POST("get_user_list", v1.GetUserList)      //  分页获取用户列表
	}
}
