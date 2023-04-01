package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(module string) *zap.SugaredLogger {
	logger, _ := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       true,
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     map[string]interface{}{"module": module},
		DisableCaller:     true,
		DisableStacktrace: true,
	}.Build()
	return logger.Sugar()
}
