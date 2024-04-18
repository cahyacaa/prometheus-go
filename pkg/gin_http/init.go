package gin_http

import (
	"github.com/gin-gonic/gin"
)

type RegisterRouter func(path string, router *gin.RouterGroup)

func InitGin() *gin.Engine {
	return gin.Default()
}
