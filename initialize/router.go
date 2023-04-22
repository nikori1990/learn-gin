package initialize

import (
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	"learn-gin/router"
)

var loginRouter = new(router.LoginRouter)
var menuRouter = new(router.MenuRouter)
var userRouter = new(router.UserRouter)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.JWTAuth())

	group := r.Group("/api")
	{
		loginRouter.Init(group)
		userRouter.Init(group)
		menuRouter.Init(group)
	}

	return r
}
