package system

import (
	"learn-gin/model/base"
)

type Permission struct {
	base.Model
	Path   string // url /api/users  /api/users/:id
	Method string // GET DELETE
}
