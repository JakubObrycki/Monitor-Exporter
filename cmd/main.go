package main

import (
	"cpu/internal/monitor"
	"fmt"
)

func main() {

	err := monitor.MonitorCpu()
	if err != nil {
		fmt.Println("Erro", err)
	}
	monitor.MonitorMem()
	monitor.LoadAverage()
	monitor.TempCpu()
}
