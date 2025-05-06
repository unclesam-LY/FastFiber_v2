package main

import (
	"FastFiber_v2/core"
	"FastFiber_v2/flags"
	"FastFiber_v2/global"
	"FastFiber_v2/routers"
	"go.uber.org/zap"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConfig()
	global.Log = core.InitZapLogger()
	// 使用zap全局替换（可选）
	zap.ReplaceGlobals(global.Log)
	global.DB = core.InitGorm()
	flags.Run()

	routers.Run()
}
