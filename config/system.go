package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // 端口值
	DbType        string `mapstructure:"tables-type" json:"dbType" yaml:"tables-type"`              // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截

}
