package middlewares

import (
	"github.com/cahyacaa/prometheus-go/pkg/prometheus_monitoring"
	"github.com/gin-gonic/gin"
	"runtime"
)

func TrackMemoryUsage(c *gin.Context) {
	var startMem runtime.MemStats
	runtime.ReadMemStats(&startMem)

	c.Next()

	var endMem runtime.MemStats
	runtime.ReadMemStats(&endMem)
	prometheus_monitoring.MemoryUsage.Set(float64(endMem.Alloc - startMem.Alloc))
}
