package loggers

import (
	"fmt"
	"hardhat-backend/config"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(GetLogger),
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger
func GetLogger(env *config.Env) Logger {
	if globalLogger == nil {
		logger := newLogger(env)
		globalLogger = &logger
	}
	return *globalLogger
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// newLogger sets up logger
func newLogger(env *config.Env) Logger {
	config := zap.NewDevelopmentConfig()
	logOutput := os.Getenv("LOG_OUTPUT")

	if env.Environment == "development" {
		fmt.Println("encode level")
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env.Environment == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	logLevel := os.Getenv("LOG_LEVEL")
	level := zap.PanicLevel
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	config.Level.SetLevel(level)

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

func GetZapCore() zapcore.Core {
	return zapLogger.Core()
}
