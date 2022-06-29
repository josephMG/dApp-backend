package loggers

import "go.uber.org/zap"

type GinLogger struct {
	*Logger
}

// GetGinLogger get the gin logger
func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// Write interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
