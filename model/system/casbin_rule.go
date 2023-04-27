package system

import (
	"learn-gin/model/base"
)

type CasbinRule struct {
	base.Model
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}
