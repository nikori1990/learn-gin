package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn-gin/global"
)

func Db() *gorm.DB {
	dsn := global.CONFIG.Mysql.Dsn()
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
