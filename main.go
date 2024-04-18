package main

import (
	"github.com/cahyacaa/prometheus-go/pkg/gin_http"
	"github.com/cahyacaa/prometheus-go/routers"
)

func main() {
	ginEng := gin_http.InitGin()
	routers.InitRouter(ginEng)
	ginEng.Run(":9000")
}
