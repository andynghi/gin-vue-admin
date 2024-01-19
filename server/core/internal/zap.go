package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Zap = new(_zap)

type _zap struct{}

// GetEncoder gets zapcore.Encoder
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

// GetEncoderConfig gets zapcore.EncoderConfig
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.GVA_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    global.GVA_CONFIG.Zap.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetEncoderCore gets the Encoder's zapcore.Core
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := FileRotatelogs.GetWriteSyncer(l.String()) // Log splitting
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder customizes log output time format
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(global.GVA_CONFIG.Zap.Prefix + t.Format("2006/01/02 - 15:04:05.000"))
}

// GetZapCores gets []zapcore.Core according to the Level of the configuration file
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := global.GVA_CONFIG.Zap.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority gets zap.LevelEnablerFunc based on zapcore.Level
// Author [SliverHorn](https://github.com/SliverHorn)
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // Debug level
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // Log level
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // warning level
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // Error level
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic level
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { //panic level
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // Terminate level
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // Debug level
			return level == zap.DebugLevel
		}
	}
}
