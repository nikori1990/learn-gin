package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
)

type TenantService struct {
}

func (TenantService) Create(c *gin.Context) {
	var tenant system.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		panic(err)
	}

	createId := tenantRepository.Create(&tenant)
	api.Success(c, createId)
}
