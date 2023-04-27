package config

type Config struct {
	Server   Server
	Mysql    Mysql
	Security Security
	Zap      Zap
}

func (config *Config) Init() {
	config.Mysql.Default()
}
