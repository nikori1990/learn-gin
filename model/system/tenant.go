package system

import (
	"learn-gin/enum"
	"learn-gin/model/base"
)

type Tenant struct {
	base.Model
	Name      string          `json:"name"`
	Code      string          `json:"code"`
	StartTime *base.LocalTime `json:"startTime"`
	EndTime   *base.LocalTime `json:"endTime"`
	Status    enum.Status     `json:"status"`
}
