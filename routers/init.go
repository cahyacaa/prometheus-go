package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("v1")
	{
		RegisterMonitoring(api)
	}
}
