package config

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt signature
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // Expiration time
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // Buffer time
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // Issuer
}
