package main

import (
	"goClass/core"
	"goClass/global"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	core.RunWindowsServer()
}
