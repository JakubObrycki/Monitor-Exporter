package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HttpListener() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
