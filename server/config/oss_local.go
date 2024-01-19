package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // Local file access path
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // Local file storage path
}
