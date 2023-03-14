package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"learn-gin/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		loadConfig(v)
	})

	loadConfig(v)
}

func loadConfig(v *viper.Viper) {
	setConfigDefault()
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}

func setConfigDefault() {
	global.CONFIG.Mysql.Default()
}
