package main

import (
	"cpu/internal/monitor"
	"cpu/internal/server"
	"time"
)

func main() {
	go func() {
		for {
			monitor.RecordMetrics()
			time.Sleep(5 * time.Second)
		}
	}()
	server.HttpListener()
}
