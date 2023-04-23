package system

import (
	"learn-gin/model/base"
)

type UserRole struct {
	base.Model
	UserId string
	RoleId string
}
