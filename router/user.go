package router

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("test")
	{
		router.GET("/test")
	}
}
