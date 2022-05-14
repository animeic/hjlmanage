package models

type Mongodb struct {
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
