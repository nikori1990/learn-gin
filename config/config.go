package config

type Config struct {
	Server   Server
	Mysql    Mysql
	Security Security
}

func (config *Config) Init() {
	config.Mysql.Default()
}
