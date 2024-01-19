package config

type DsnProvider interface {
	Dsn() string
}

// The Embedded structure can be flattened to the previous layer, thus keeping the structure of the config file the same as before.
// See playground: https://go.dev/play/p/KIcuhqEoxmY

// GeneralDB is also used unchanged by Pgsql and Mysql
type GeneralDB struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`       // Advanced configuration
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // Database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // Database password
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // Database password
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //Database engine, default InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // Whether to enable Gorm global log
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // The maximum number of idle connections
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // Maximum number of connections opened to the database
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //Whether to enable global disabling of plural numbers, true means enabled
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // Whether to write log files through zap
}

type SpecializedDB struct {
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	Disable   bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}
