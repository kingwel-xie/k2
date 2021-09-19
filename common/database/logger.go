package database

import (
	"context"
	"fmt"
	"time"

	"github.com/kingwel-xie/k2/core/logger"

	glogger "gorm.io/gorm/logger"
)

var log = logger.Logger("gorm")

type gormLogger struct {
	SlowThreshold time.Duration
}

// LogMode log mode
func (l *gormLogger) LogMode(level glogger.LogLevel) glogger.Interface {
	newlogger := *l
	return &newlogger
}

// Info print info
func (l gormLogger) Info(_ctx context.Context, msg string, data ...interface{}) {
	log.Infof(msg, data...)
}

// Warn print warn messages
func (l gormLogger) Warn(_ctx context.Context, msg string, data ...interface{}) {
	log.Warnf(msg, data...)
}

// Error print error messages
func (l gormLogger) Error(_ctx context.Context, msg string, data ...interface{}) {
	log.Errorf(msg, data...)
}

// Trace print sql message
func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil:
		sql, rows := fc()
		log.Errorf("SQL: %v, elapsed=%f, %s, rows=%d", err, float64(elapsed.Nanoseconds())/1e6, sql, rows)

	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		log.Warnf("SQL: %s, elapsed=%f, %s, rows=%d", slowLog, float64(elapsed.Nanoseconds())/1e6, sql, rows)

	default:
		sql, rows := fc()
		log.Debugf("SQL: elapsed=%f, %s, rows=%d", float64(elapsed.Nanoseconds())/1e6, sql, rows)
	}
}
