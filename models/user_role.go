package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId string
	RoleId string
}
