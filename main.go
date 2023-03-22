package main

import (
	"learn-gin/core"
	"learn-gin/router"
)

func main() {
	core.Viper()
	core.Db()

	router := router.SetUpRouters()
	routerErr := router.Run(":8000")
	if routerErr != nil {
		return
	}
}
