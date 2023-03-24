package models

type Menu struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Path     string  `json:"path"` // /dashboard /system/users
	Icon     string  `json:"icon"`
	Pid      string  `json:"pid"`
	Children []*Menu `json:"children" gorm:"-"`
}

// TableName 表示配置操作数据库的表名
func (Menu) TableName() string {
	return "menu"
}
