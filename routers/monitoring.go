package routers

import (
	"github.com/cahyacaa/prometheus-go/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMonitoring(rg *gin.RouterGroup) {
	api := rg.Group("/monitoring")
	{
		api.GET("", controllers.CollectData)
	}
}
