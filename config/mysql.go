package config

type Mysql struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Config   string
}

func (m *Mysql) Default() {
	m.Host = "127.0.0.1"
	m.Port = "3306"
	m.Config = "charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}
