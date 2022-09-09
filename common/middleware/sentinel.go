package middleware

import (
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/alibaba/sentinel-golang/logging"
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/gin-gonic/gin"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	_ = logging.ResetGlobalLogger(log.WithCallerSkip(1))
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
