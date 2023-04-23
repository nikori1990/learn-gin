package system

import (
	"learn-gin/model/base"
)

type Dept struct {
	base.Model
	Name     string `json:"name"`
	ParentId uint   `json:"parentId"`
}
