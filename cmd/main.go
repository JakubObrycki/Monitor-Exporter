package main

import (
	"cpu/internal/monitor"
)

func main() {
	monitor.MonitorCpu()
	monitor.MonitorMem()
	monitor.LoadAverage()
}
