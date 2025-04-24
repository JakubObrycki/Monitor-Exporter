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
		fmt.Println("Error", err) // pozmieniac wszedzie tu na logi, ostatecznie do wywalenie wszystkie fmt.Prity z totalmemory itp
	}
	fmt.Println("Available Memory:", dataMem.Available, "\nUsed Memory:", data.Used, "\nTotal Memory:", dataMem.TotalMemory, "\nUsedPercent Memory:", dataMem.UsedPercent, "\nFree Memory:", dataMem.Free)

	dataLa, err := monitor.LoadAverage()
	if err != nil {
		fmt.Print("Error", err)
	}
	fmt.Println("Load Average", dataLa.Load1, dataLa.Load5, dataLa.Load15)

	monitor.TempCpu() // zmiana funkcje nie czujniki temperatury tylko cos innego

	dataCpuTemp, err := monitor.SystemUpTime()
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("System time:", dataCpuTemp.Days, "days", dataCpuTemp.Hours, "hours", dataCpuTemp.Minutes, "minutes")
}
