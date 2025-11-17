package metrics

import "github.com/prometheus/client_golang/prometheus"

func register() {
	prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
}
