package initialize

import (
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/middlewares"
	"learn-gin/router"
	"learn-gin/router/system"
)

var tenantRouter = new(system.TenantRouter)
var deptRouter = new(system.DeptRouter)
var roleRouter = new(system.RoleRouter)
var permissionRouter = new(system.PermissionRouter)
var userRouter = new(system.UserRouter)
var userRoleRouter = new(system.UserRoleRouter)
var rolePermissionRouter = new(system.RolePermissionRouter)
var menuRouter = new(system.MenuRouter)

var loginRouter = new(router.LoginRouter)

func InitRouters() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Recover())
	r.Use(middlewares.JWTAuth())
	r.Use(middlewares.CasbinHandler())

	apiRouter := r.Group("/api")
	{
		systemRouter := system.InitSystemRouter(apiRouter)
		{
			tenantRouter.Init(systemRouter)
			userRouter.Init(systemRouter)
			menuRouter.Init(systemRouter)
			deptRouter.Init(systemRouter)
			roleRouter.Init(systemRouter)
			permissionRouter.Init(systemRouter)
			userRoleRouter.Init(systemRouter)
			rolePermissionRouter.Init(systemRouter)
		}

		loginRouter.Init(apiRouter)
	}

	global.Logger.Info("initialize router success")
	return r
}
