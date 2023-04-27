package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/model/api"
	"learn-gin/utils"
	"strconv"
)

type LoginForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type LoginService struct {
}

func (s LoginService) Login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		fmt.Println(err)
	}

	user := userRepository.GetByName(form.Username)
	if user.Password != form.Password {
		global.Logger.Error("invalid password")
		panic("invalid password")
	}

	userRoles := userRoleRepository.ListByUserId(user.ID)
	var roleIds []uint
	for _, role := range userRoles {
		roleIds = append(roleIds, role.ID)
	}

	var permissionsIds []uint
	if roleIds != nil {
		rolePermissions := rolePermissionRepository.ListByRoleIds(roleIds)
		if len(rolePermissions) > 0 {
			for _, permission := range rolePermissions {
				permissionsIds = append(permissionsIds, permission.PermissionId)
			}
		}
	}

	var resources [][]string
	if permissionsIds != nil {
		permissions := permissionRepository.ListByIds(permissionsIds)
		if len(permissions) > 0 {
			for _, permission := range permissions {
				resources = append(resources, []string{permission.Path, permission.Method})
			}
		}
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(user.TenantId)), user.Username, resources)

	if err != nil {
		panic("GenerateToken error")
	}

	global.Logger.Infof("token: %s", token)

	api.Success(c, gin.H{
		"token": token,
	})
}
