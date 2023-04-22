package system

type SysMenu struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"` // /dashboard /system/users
	Icon     string     `json:"icon"`
	Pid      string     `json:"pid"`
	Children []*SysMenu `json:"children" gorm:"-"`
}

// TableName 表示配置操作数据库的表名
func (SysMenu) TableName() string {
	return "menu"
}
