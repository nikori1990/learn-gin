package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string // admin operator guest
}
