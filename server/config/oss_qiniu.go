package config

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                  // Storage area
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                            // Space name
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"img-path"`                      // CDN accelerated domain name
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`                // Secret key AK
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`                // Secret key SK
	UseHTTPS      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`                   // Whether to use https
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"` // Whether to use CDN upload acceleration for uploading
}
