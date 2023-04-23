package initialize

import (
	"fmt"
	"learn-gin/core"
	"learn-gin/global"
)

func InitConfig() {
	global.CONFIG.Init()

	fmt.Printf("config %v", global.CONFIG)

	core.Viper()
}
