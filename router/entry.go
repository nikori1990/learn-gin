package router

import "learn-gin/service/system"

var (
	deptService           = new(system.DeptService)
	menuService           = new(system.MenuService)
	permissionService     = new(system.PermissionService)
	rolePermissionService = new(system.RolePermissionService)
	roleService           = new(system.RoleService)
	userRoleService       = new(system.UserRoleService)
	userService           = new(system.UserService)
)
