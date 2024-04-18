package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CollectData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("OK at %s", time.Now().String()),
	})
	return
}
