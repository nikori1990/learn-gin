package routers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	v1 "learn-gin/routers/v1"
)

func SetUpRouters() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.JWTAuth())

	apiRouter := router.Group("/api/v1")
	v1.LoginRoutersInit(apiRouter)
	v1.MenuRoutersInit(apiRouter)
	v1.UserRoutersInit(apiRouter)

	return router
}
