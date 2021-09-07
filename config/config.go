package config

type Server struct {
	JWT      JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System   System   `mapstructure:"system" json:"system" yaml:"system"`
	Zap      Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql    Mysql    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	AutoCode Autocode `mapstructure:"autoCode" json:"autoCode" yaml:"autoCode"`
}
