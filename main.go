package main

import (
	"witcier/go-api/core"
	"witcier/go-api/global"
	"witcier/go-api/initialize"

	"go.uber.org/zap"
)

func main() {
	global.Viper = core.Viper()
	global.Log = core.Zap()
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm()

	core.RunServer()
}
