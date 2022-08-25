package config

type Mysql struct {
	Path         string `mapstructure:"path" json:"path"`
	Port         string `mapstructure:"port" json:"port"`
	Prefix       string `mapstructure:"prefix" json:"prefix"`
	Config       string `mapstructure:"config" json:"config"`
	DbName       string `mapstructure:"db-name" json:"db-name"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
