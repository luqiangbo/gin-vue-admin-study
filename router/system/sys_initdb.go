package system

import (
	"github.com/gin-gonic/gin"
	"gin-vue-admin-study/api/v1"
)

type InitRouter struct {
}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	var dbApi = v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("init_db", dbApi.InitDB)
	}
}
