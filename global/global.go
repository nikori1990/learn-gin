package global

import (
	"github.com/spf13/viper"
	"learn-gin/config"
)

var (
	VIPER  *viper.Viper
	CONFIG config.Config
)
