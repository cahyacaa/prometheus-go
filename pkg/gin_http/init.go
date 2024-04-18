package gin_http

import (
	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	return gin.Default()
}
