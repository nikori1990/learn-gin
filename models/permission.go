package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Path   string // url /api/users  /api/users/:id
	Method string // GET DELETE
}
