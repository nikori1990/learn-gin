package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
)

type UserRoleService struct {
}

func (UserRoleService) Create(c *gin.Context) {
	var userRole system.UserRole
	if err := c.ShouldBindJSON(&userRole); err != nil {
		panic(err)
	}
	createId := userRoleRepository.Create(&userRole)
	api.Success(c, createId)
}

func (UserRoleService) List(c *gin.Context) {
	userRoleList := userRoleRepository.List()

	var userIds []uint
	var roleIds []uint
	for _, userRole := range userRoleList {
		userIds = append(userIds, userRole.UserId)
		roleIds = append(roleIds, userRole.RoleId)
	}

	roleList := roleRepository.ListByIds(roleIds)
	userList := userRepository.ListByIds(userIds)

	roleMap := make(map[uint]*system.Role)
	userMap := make(map[uint]*system.User)

	for _, role := range roleList {
		roleMap[role.ID] = role
	}

	for _, user := range userList {
		userMap[user.ID] = user
	}

	for _, userRole := range userRoleList {
		user := userMap[userRole.UserId]
		role := roleMap[userRole.RoleId]

		userRole.Username = user.Username
		userRole.RoleName = role.Name
	}

	api.Success(c, userRoleList)
}
