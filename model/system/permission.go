package system

import (
	"learn-gin/model/base"
)

type Permission struct {
	base.Model
	Path   string `json:"path"`   // url /api/users  /api/users/:id
	Method string `json:"method"` // GET DELETE
}
