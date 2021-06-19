package core

import "goClass/initialize"

func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Run()
}
