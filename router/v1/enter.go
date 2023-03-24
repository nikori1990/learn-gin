package v1

import "github.com/gin-gonic/gin"

type v1Router struct {
	LoginRouter
	MenuRouter
	UserRouter
}

func SetUpV1Router(router *gin.Engine) {
	var v1Router = new(v1Router)
	loginRouter := v1Router.LoginRouter
	userRouter := v1Router.UserRouter
	menuRouter := v1Router.MenuRouter

	v1Group := router.Group("/api/v1")
	{
		loginRouter.Init(v1Group)
		userRouter.Init(v1Group)
		menuRouter.Init(v1Group)
	}
}
