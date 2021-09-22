package middleware

import (
	sentinelPlugin "github.com/alibaba/sentinel-golang/adapter/gin"
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/alibaba/sentinel-golang/logging"
	"github.com/gin-gonic/gin"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	logging.ResetGlobalLogger(log)
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
