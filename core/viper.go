package core

import (
	"fmt"
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

	v.SetDefault("env", "")

	loadConfig(v)

	env := v.GetString("env")

	if env != "" {
		loadEnvConfig(v, env)
	}
}

func loadConfig(v *viper.Viper) {
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}

func loadEnvConfig(v *viper.Viper, env string) {
	fileName := fmt.Sprintf("config-%s.yaml", env)

	v.SetConfigFile(fileName)

	err := v.MergeInConfig()
	if err != nil {
		// 如果环境配置文件不存在，则忽略错误
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
}
