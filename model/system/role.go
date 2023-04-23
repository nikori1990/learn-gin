package system

import "learn-gin/model/base"

type Role struct {
	base.Model
	Name string `json:"name" gorm:"name"` // admin operator guest
}
