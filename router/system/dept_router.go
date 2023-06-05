package system

import (
	"github.com/gin-gonic/gin"
)

type DeptRouter struct {
}

func (DeptRouter) Init(group *gin.RouterGroup) {
	router := group.Group("/depts")
	{
		router.GET("", deptService.List)
		router.GET("/:id", deptService.GetById)
		router.POST("", deptService.Create)
		router.PUT("", deptService.Update)
		router.DELETE("/:id", deptService.Delete)
		router.GET("/page", deptService.Page)
		router.GET("/tree", deptService.Tree)
	}
}
