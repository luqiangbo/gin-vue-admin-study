package core

import (
	"fmt"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/initialize"
)

func RunWindowsServer() {

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	Router.Run(address)
}
