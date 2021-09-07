package config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer_restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server_api" yaml:"server-api"`
	SInitialize     string `mapstructure:"server-initialize" json:"server_initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server_model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server_request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server_router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server_service" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web_api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web_form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web_table" yaml:"web-table"`
	WFlow           string `mapstructure:"web-flow" json:"web_flow" yaml:"web-flow"`
}
