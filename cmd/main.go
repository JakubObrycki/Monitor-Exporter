package main

import (
	"cpu/internal/monitor"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		for {
			monitor.RecordMetrics()
			time.Sleep(5 * time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
