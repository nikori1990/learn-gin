package v1

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	LoginRouter
	MenuRouter
	UserRouter
}

func (v1 *Router) Init(router *gin.Engine) {
	loginRouter := v1.LoginRouter
	userRouter := v1.UserRouter
	menuRouter := v1.MenuRouter

	v1Group := router.Group("/api/v1")
	{
		loginRouter.Init(v1Group)
		userRouter.Init(v1Group)
		menuRouter.Init(v1Group)
	}
}
