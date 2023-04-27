package initialize

import (
	"learn-gin/core"
	"learn-gin/global"
)

func InitLogger() {
	logger := core.InitLogger()
	global.Logger = logger
}
