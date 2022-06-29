package loggers

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/inconshreveable/log15"
	"go.uber.org/zap"
)

type swapHandler struct {
	handler atomic.Value
}

type EtherLogger struct {
	*Logger
	ctx []interface{}
	h   *swapHandler
}

// GetEtherLogger gets logger for go-ethereum
func (l *Logger) GetEtherLogger() *EtherLogger {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	return &EtherLogger{Logger: newSugaredLogger(logger)}
}

// Ethereum Framework Logger (log15) Interface Implementations
// ---- START ----
func (l *EtherLogger) New(ctx ...interface{}) log15.Logger {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(2),
	)
	child := &EtherLogger{
		Logger: newSugaredLogger(logger),
	}
	child.SetHandler(l.h.handler.Load().(log15.Handler))
	return child
}

func (l *EtherLogger) Trace(msg string, ctx ...interface{}) {
	core := GetZapCore()
	elapsed := time.Now()
	if core.Enabled(zap.DebugLevel) {
		l.Debug("[", elapsed.UnixMilli(), " ms ] ", msg)
	}

	if core.Enabled(zap.WarnLevel) {
		l.SugaredLogger.Warn("[", elapsed.UnixMilli(), " ms ] ", msg)
	}

	if core.Enabled(zap.ErrorLevel) {
		l.SugaredLogger.Error("[", elapsed.UnixMilli(), " ms ] ", msg)
	}
}

func (l *EtherLogger) Debug(msg string, ctx ...interface{}) {
	l.Debugf(msg)
}

func (l *EtherLogger) Info(msg string, ctx ...interface{}) {
	l.Infof(msg)
}

func (l *EtherLogger) Warn(msg string, ctx ...interface{}) {
	l.Warnf(msg)
}

func (l *EtherLogger) Error(msg string, ctx ...interface{}) {
	l.Errorf(msg)
}

func (l *EtherLogger) Crit(msg string, ctx ...interface{}) {
	l.Fatalf(msg)
	os.Exit(1)
}

func (l *EtherLogger) GetHandler() log15.Handler {
	return l.h.handler.Load().(log15.Handler)
}

func (l *EtherLogger) SetHandler(h log15.Handler) {
}
