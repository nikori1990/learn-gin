package system

import (
	"learn-gin/repository/system"
)

var (
	tenantRepository         = new(system.TenantRepository)
	deptRepository           = new(system.DeptRepository)
	menuRepository           = new(system.MenuRepository)
	permissionRepository     = new(system.PermissionRepository)
	roleRepository           = new(system.RoleRepository)
	userRoleRepository       = new(system.UserRoleRepository)
	userRepository           = new(system.UserRepository)
	rolePermissionRepository = new(system.RolePermissionRepository)
)
