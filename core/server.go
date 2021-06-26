package core

import (
	"fmt"
	"goClass/global"
	"goClass/initialize"
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
