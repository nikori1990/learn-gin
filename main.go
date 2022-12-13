package main

import (
	"fmt"
	"github.com/spf13/viper"
	"learn-gin/routers"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	configErr := viper.ReadInConfig()
	if configErr != nil {
		panic(fmt.Errorf("fatal error config file: %w", configErr))
	}

	fmt.Println("name", viper.Get("name"))

	router := routers.SetUpRouters()
	routerErr := router.Run(":8000")
	if routerErr != nil {
		return
	}
}
