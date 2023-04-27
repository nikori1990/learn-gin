package config

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`           // 级别
	LogDirectory string `mapstructure:"director" json:"director"  yaml:"director"` // 日志文件夹
	MaxSize      int    `mapstructure:"maxSize" json:"maxSize" yaml:"maxSize"`
	MaxBackups   int    `mapstructure:"maxBackups" json:"maxBackups" yaml:"maxBackups"`
	MaxAge       int    `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"` // 日志留存时间
}
