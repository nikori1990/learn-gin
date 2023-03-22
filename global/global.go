package global

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"learn-gin/config"
)

var (
	VIPER  *viper.Viper
	CONFIG config.Config
	DB     *gorm.DB
)
