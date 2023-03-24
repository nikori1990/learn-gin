package main

import (
	"learn-gin/core"
	"learn-gin/global"
	"learn-gin/router"
)

func main() {
	core.Viper()
	global.DB = core.Db()

	router := router.SetUpRouters()
	routerErr := router.Run(":8000")
	if routerErr != nil {
		return
	}
}
