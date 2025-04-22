package main

import (
	"cpu/internal/monitor"
	"fmt"
)

func main() {

	data, err := monitor.MonitorCpu()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Total CPU:", data.CPU, "%")

	monitor.MonitorMem()
	monitor.LoadAverage()
	monitor.TempCpu()

	dataCpu, err := monitor.SystemUpTime()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("\nSystem time:", dataCpu.Days, dataCpu.Hours, dataCpu.Minutes)
}
