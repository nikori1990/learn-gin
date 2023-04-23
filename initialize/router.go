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
var permissionRouter = new(router.PermissionRouter)
var userRoleRouter = new(router.UserRoleRouter)
var rolePermissionRouter = new(router.RolePermissionRouter)

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
		permissionRouter.Init(apiRouter)
		userRoleRouter.Init(apiRouter)
		rolePermissionRouter.Init(apiRouter)
	}

	return r
}
