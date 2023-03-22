package models

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleId       string
	PermissionId string
}
