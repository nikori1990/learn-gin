package query

import (
	"learn-gin/model/base"
)

type UserQuery struct {
	base.PageQuery
	Name string `json:"name" form:"name"`
}
