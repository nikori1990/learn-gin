package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service"
)

type DeptRouter struct {
}

var deptService = new(service.DeptService)

func (DeptRouter) Init(group *gin.RouterGroup) {
	deptRouter := group.Group("/depts")
	{
		deptRouter.GET("", deptService.List)
		deptRouter.GET("/:id", deptService.GetById)
		deptRouter.POST("", deptService.Create)
		deptRouter.PUT("", deptService.Update)
		deptRouter.DELETE("/:id", deptService.Delete)
	}
}
