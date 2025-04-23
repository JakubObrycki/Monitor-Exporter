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

	dataMem, err := monitor.MonitorMem()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("\n--Memory--", "\nAvailable Memory:", dataMem.Available, "\nUsed Memory:", data.Used, "\nTotal Memory:", dataMem.TotalMemory, "\nUsedPercent Memory:", dataMem.UsedPercent, "\nFree Memory:", dataMem.Free)

	monitor.LoadAverage()
	monitor.TempCpu()

	dataCpuTemp, err := monitor.SystemUpTime()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("\nSystem time:", dataCpuTemp.Days, "days", dataCpuTemp.Hours, "hours", dataCpuTemp.Minutes, "minutes")
}
