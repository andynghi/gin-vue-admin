package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // Recipients: Multiple separated by English commas. Example: a@qq.com b@qq.com Please put this project under formal development Use as parameter
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // Sender The email address you want to send the email to
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // Server address For example smtp.qq.com Please go to QQ or the email address you want to send the email to check its smtp protocol
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // Key The key used for login. It is best not to use the email password. Go to the email SMTP to apply for a key used for login.
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // Nickname The sender's nickname is usually your own email address
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // Port Please go to QQ or the email address you want to send the email to check its SMTP protocol. Most of them are 465.
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // Whether to SSL or not. Whether to enable SSL
}
