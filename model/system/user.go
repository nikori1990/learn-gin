package system

import (
	"learn-gin/model/base"
)

type User struct {
	base.Model
	Username string `json:"username"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
