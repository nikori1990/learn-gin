package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	v1 "learn-gin/router/v1"
)

func SetUpRouters() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.JWTAuth())

	v1.SetUpV1Router(router)

	return router
}
