package system

import (
	"github.com/gin-gonic/gin"
	v1 "gin-vue-admin-study/api/v1"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	userRouter := Router.Group("user")
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)       // 注册
		userRouter.POST("register", baseApi.Register) // 用户注册账号
	}
	return baseRouter
}
