package initialize

import (
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/middlewares"
	"learn-gin/router"
)

var tenantRouter = new(router.TenantRouter)
var deptRouter = new(router.DeptRouter)
var roleRouter = new(router.RoleRouter)
var permissionRouter = new(router.PermissionRouter)
var userRouter = new(router.UserRouter)
var userRoleRouter = new(router.UserRoleRouter)
var rolePermissionRouter = new(router.RolePermissionRouter)
var menuRouter = new(router.MenuRouter)

var loginRouter = new(router.LoginRouter)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Recover())
	r.Use(middlewares.JWTAuth())
	r.Use(middlewares.CasbinHandler())

	apiRouter := r.Group("/api")
	{
		tenantRouter.Init(apiRouter)
		userRouter.Init(apiRouter)
		menuRouter.Init(apiRouter)
		deptRouter.Init(apiRouter)
		roleRouter.Init(apiRouter)
		permissionRouter.Init(apiRouter)
		userRoleRouter.Init(apiRouter)
		rolePermissionRouter.Init(apiRouter)

		loginRouter.Init(apiRouter)
	}

	global.Logger.Info("initialize router success")
	return r
}
