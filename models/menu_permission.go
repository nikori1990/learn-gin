package models

import "gorm.io/gorm"

type MenuPermission struct {
	gorm.Model
	MenuId       string
	PermissionId string
}
