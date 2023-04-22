package global

import (
	"gorm.io/gorm"
	"learn-gin/config"
)

var (
	CONFIG config.Config
	DB     *gorm.DB
)
