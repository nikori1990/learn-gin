package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

// TableName 表示配置操作数据库的表名
func (User) TableName() string {
	return "user"
}
