package initialize

import (
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	"learn-gin/router"
)

var loginRouter = new(router.LoginRouter)
var menuRouter = new(router.MenuRouter)
var userRouter = new(router.UserRouter)
var deptRouter = new(router.DeptRouter)
var roleRouter = new(router.RoleRouter)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.JWTAuth())

	apiRouter := r.Group("/api")
	{
		loginRouter.Init(apiRouter)
		userRouter.Init(apiRouter)
		menuRouter.Init(apiRouter)
		deptRouter.Init(apiRouter)
		roleRouter.Init(apiRouter)
	}

	return r
}
