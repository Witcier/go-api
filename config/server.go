package config

type Server struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
	Cors   CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
