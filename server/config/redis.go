package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // Server address: port
	Password string `mapstructure:"password" json:"password" yaml:"password"` // Password
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // Which database of redis
}
