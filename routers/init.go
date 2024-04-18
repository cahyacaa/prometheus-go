package routers

import (
	"github.com/cahyacaa/prometheus-go/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("v1")
	{
		api.Use(middlewares.TrackMemoryUsage)
		RegisterMonitoring(api)
		api.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}
