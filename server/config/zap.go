package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // Level
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // Log prefix
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // Output
	Director      string `mapstructure:"director" json:"director" yaml:"director"`                   // Log folder
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // Encoding level
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // stack name

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`                      // Log retention time
	ShowLine     bool `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // Show line
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // Output console
}

// ZapEncodeLevel returns zapcore.LevelEncoder based on EncodeLevel
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // Lowercase encoder (default)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // Lowercase encoder with color
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // Capital encoder
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // Capital encoder with color
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel is converted to zapcore.Level according to the string
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *Zap) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
