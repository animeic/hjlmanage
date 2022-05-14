package models

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`
	Mode    string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Port    string `mapstructure:"port" json:"port" yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
}
