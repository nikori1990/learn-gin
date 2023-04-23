package system

import "learn-gin/repository"

var (
	deptRepository           = new(repository.DeptRepository)
	menuRepository           = new(repository.MenuRepository)
	permissionRepository     = new(repository.PermissionRepository)
	roleRepository           = new(repository.RoleRepository)
	userRoleRepository       = new(repository.UserRoleRepository)
	userRepository           = new(repository.UserRepository)
	rolePermissionRepository = new(repository.RolePermissionRepository)
)
