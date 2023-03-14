package main

import (
	"fmt"
	"learn-gin/core"
	"learn-gin/global"
	"learn-gin/routers"
)

func main() {
	core.Viper()
	fmt.Println(global.CONFIG.Mysql.Dns())

	router := routers.SetUpRouters()
	routerErr := router.Run(":8000")
	if routerErr != nil {
		return
	}
}
