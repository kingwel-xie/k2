// Package log is the logging library used by IPFS & libp2p
// (https://github.com/ipfs/go-ipfs).
package logger

import (
	"time"

	"go.uber.org/zap"
)

var (
	// DefaultLogger logger
	DefaultLogger *KLogger
)

// StandardLogger provides API compatibility with standard printf loggers
// eg. go-logging
type StandardLogger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
}

// Logger retrieves an event logger by name
func logger(system string) *KLogger {
	if len(system) == 0 {
		panic("missing name")
	}

	logger := getLogger(system)
	return &KLogger{
		system:        system,
		SugaredLogger: *logger,
	}
}

// KLogger implements the StandardLogger interface
type KLogger struct {
	zap.SugaredLogger
	system     string
}

// FormatRFC3339 returns the given time in UTC with RFC3999Nano format.
func FormatRFC3339(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// Logger retrieves an event logger by name
func Logger(system string) *KLogger {
	return logger(system)
}

func (h *KLogger) WithError(err error) *KLogger {
	h.SugaredLogger = *h.SugaredLogger.With("error", err)
	return h
}

func (h *KLogger) WithFields(args ...interface{}) *KLogger {
	h.SugaredLogger = *h.SugaredLogger.With(args...)
	return h
}

func (h *KLogger) WithCallerSkip(skip int) *KLogger {
	h.SugaredLogger = *h.SugaredLogger.Desugar().WithOptions(zap.AddCallerSkip(skip)).Sugar()
	return h
}

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
