package middleware

import (
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/alibaba/sentinel-golang/logging"
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
	"github.com/kingwel-xie/k2/core/logger"
)

type logAdapter struct {
	*logger.KLogger
}

func (l logAdapter) Debug(msg string, keysAndValues ...interface{}) {
	l.KLogger.Debugw(msg, keysAndValues...)
}

func (l logAdapter) DebugEnabled() bool {
	return true
}

func (l logAdapter) Info(msg string, keysAndValues ...interface{}) {
	l.KLogger.Infow(msg, keysAndValues...)
}

func (l logAdapter) InfoEnabled() bool {
	return true
}

func (l logAdapter) Warn(msg string, keysAndValues ...interface{}) {
	l.KLogger.Warnw(msg, keysAndValues...)
}

func (l logAdapter) WarnEnabled() bool {
	return true
}

func (l logAdapter) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "Error", err)
	l.KLogger.Errorw(msg, keysAndValues...)
}

func (l logAdapter) ErrorEnabled() bool {
	return true
}

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	_ = logging.ResetGlobalLogger(logAdapter{ log.WithCallerSkip(1) })
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 200,
			Strategy:     system.BBR,
		},
	}); err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	return sentinelPlugin.SentinelMiddleware()
}
