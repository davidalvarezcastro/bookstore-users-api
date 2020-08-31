package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log variable
	log *zap.Logger
)

func init() {
	var err error
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// GetLogger returns the logger var
func GetLogger() *zap.Logger {
	return log
}

// Field returns a Field from a string key
func Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// Debug logs an debug message
func Debug(msg string, tags ...zap.Field) {
	defer log.Sync()
	log.Debug(msg, tags...)
}

// Info logs an info message
func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	defer log.Sync()
}

// Error logs an error message
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	defer log.Sync()
}
