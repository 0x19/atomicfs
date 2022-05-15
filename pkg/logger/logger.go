package logger

import (
	"strings"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func determineLogLevel(verbosityLevel string) zapcore.Level {
	var zapLevel zapcore.Level
	verbosityLevel = strings.ToLower(verbosityLevel)
	switch verbosityLevel {
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "panic":
		zapLevel = zapcore.DPanicLevel
	default:
		zapLevel = zapcore.DebugLevel
	}
	return zapLevel
}

func New(level string, environment string, global bool, options []zap.Option, labels []zapcore.Field) (*Logger, error) {
	var logger *zap.Logger
	var err error

	if environment == "development" {
		config := zap.NewDevelopmentEncoderConfig()
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(colorable.NewColorableStdout()),
			determineLogLevel(level),
		))
	} else if environment == "production" {
		logger, err = zap.NewProduction(options...)
		if err != nil {
			return nil, err
		}
	}

	toReturn := &Logger{Logger: logger.With(labels...)}

	// Replace the global zap logger with built logger
	if global {
		zap.ReplaceGlobals(toReturn.Logger)
	}

	return toReturn, nil
}
