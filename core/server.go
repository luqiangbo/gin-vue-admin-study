package core

import (
	"fmt"
	"goClass/global"
	"goClass/initialize"
)

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	Router.Run(address)
}
