package router

import "github.com/gin-gonic/gin"

type TenantRouter struct {
}

func (TenantRouter) Init(group *gin.RouterGroup) {
	tenantRouter := group.Group("/tenants")
	{
		tenantRouter.POST("", tenantService.Create)
	}
}
