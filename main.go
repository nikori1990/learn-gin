package main

import (
	"fmt"
	"learn-gin/core"
	"learn-gin/global"
	"learn-gin/router"
)

func main() {

	global.CONFIG.Init()
	core.Viper()
	global.DB = core.Db()

	port := global.CONFIG.Server.Port

	fmt.Printf("config %v", global.CONFIG)

	r := router.SetUpRouters()
	routerErr := r.Run(fmt.Sprintf(":%v", port))
	if routerErr != nil {
		return
	}
}
