package main

import (
	"fmt"
	"learn-gin/global"
	"learn-gin/initialize"
)

func main() {

	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDb()
	initialize.InitCasbin()

	router := initialize.InitRouters()

	port := global.CONFIG.Server.Port
	routerErr := router.Run(fmt.Sprintf(":%v", port))
	if routerErr != nil {
		return
	}

}
