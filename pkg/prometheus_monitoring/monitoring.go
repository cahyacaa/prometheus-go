package prometheus_monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var MemoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "api_request_memory_usage_bytes",
	Help: "Memory usage of the API request",
})

func InitPrometheus() {
	prometheus.MustRegister(MemoryUsage)
}
