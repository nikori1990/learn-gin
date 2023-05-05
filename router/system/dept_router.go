package system

import (
	"github.com/gin-gonic/gin"
)

type DeptRouter struct {
}

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
