package core

import (
	"fmt"
	"net/http"
	"time"
	"witcier/go-api/global"
	"witcier/go-api/initialize"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.Config.System.UseMultipoint || global.Config.System.UseRedis {
		initialize.Redis()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)

	global.Log.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
