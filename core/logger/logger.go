package logger

import (
	logging "github.com/ipfs/go-log/v2"
)

var (
	// DefaultLogger logger
	DefaultLogger *logging.ZapEventLogger
)


type KLogger struct {
	logging.ZapEventLogger
}

// Logger retrieves an event logger by name
func Logger(system string) *KLogger {
	log := logging.Logger(system)
	return &KLogger{ZapEventLogger: *log}
}

func (h *KLogger) WithError(err error) *KLogger {
	h.ZapEventLogger.SugaredLogger = *h.ZapEventLogger.SugaredLogger.With("error", err)
	return h
}

func (h *KLogger) WithFields(args ...interface{}) *KLogger {
	h.ZapEventLogger.SugaredLogger = *h.ZapEventLogger.SugaredLogger.With(args...)
	return h
}

// default logger
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	DefaultLogger.Infof(template, args...)
}

func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	DefaultLogger.Debugf(template, args...)
}

func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	DefaultLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	DefaultLogger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

func Fatalf(template string, args ...interface{}) {
	DefaultLogger.Panicf(template, args...)
}

