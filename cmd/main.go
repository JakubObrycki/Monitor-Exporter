package main

import (
	"cpu/internal/monitor"
	"log"
)

func main() {
	data, err := monitor.MonitorCpu() // pomyslec o go routines
	if err != nil {
		log.Printf("MonitorCpu error: %v", err)
	}
	log.Println("Total CPU:", data.CPU, "%")

	dataMem, err := monitor.MonitorMem()
	if err != nil {
		log.Printf("Monitor memory error: %v", err)
	}
	log.Println("Available Memory:", dataMem.Available, "\nUsed Memory:", data.Used, "\nTotal Memory:", dataMem.TotalMemory, "\nUsedPercent Memory:", dataMem.UsedPercent, "\nFree Memory:", dataMem.Free)

	dataLa, err := monitor.LoadAverage()
	if err != nil {
		log.Printf("Load Average error: %v", err)
	}
	log.Println("Load Average", dataLa.Load1, dataLa.Load5, dataLa.Load15)

	dataCpuTemp, err := monitor.SystemUpTime()
	if err != nil {
		log.Printf("System Uptime error %v", err)
	}
	log.Println("System time:", dataCpuTemp.Days, "days", dataCpuTemp.Hours, "hours", dataCpuTemp.Minutes, "minutes")
}
