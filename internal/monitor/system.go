package monitor

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

// dodac jeszcze inne ciekawe funkcji ktore moga byc atrakcyjne w typ wypadku
type DataStat struct { // dodanie tego zeby wyprowadzilo metryki z dockera prometheusa i dodalo do grafany
	CPU    float64 `json:"CPU"`
	Memory float64 `json:"Memory"`
}

func MonitorCpu() error { // poprawic jeszcze tutaj i mem rowniez zgodnie z biblioteka // dodac alerting, ewentualnie przekminic to w grafanie zeby to grafana wysylala

	usage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Errorf("error with CPU monitoring %w", err)
	}

	fmt.Printf("Total CPU:  %.2f%%\n", usage[0])
	if usage[0] > 80.0 {
		fmt.Println("--Too high CPU!--")
	}

	return nil
}

func MonitorMem() error {

	v, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Errorf("error with memory monitoring %w", err)
	}

	totalMemory := float64(v.Total) / (1024 * 1024 * 1024)
	roundtotalmemory := math.Round(totalMemory*100) / 100

	fmt.Printf("Total Memory: %.2f GB\n", roundtotalmemory)
	if roundtotalmemory > 80.0 {
		fmt.Println("--Too high Memory!--")
	}

	return nil
}

// Load Average obciazenie systemu
func LoadAverage() error {

	lavg, err := load.Avg()
	if err != nil {
		return fmt.Errorf("eror with loadaverage %w", err)
	}

	fmt.Printf("Load Average: 1m: %.2f, \nLoad Average 5m: %.2f \nLoad Average 15m: %.2f", lavg.Load1, lavg.Load5, lavg.Load15)

	return nil
}

// temperatura CPU

func TempCpu() error { // funckcja nie dziala ?

	temp, err := host.SensorsTemperatures()
	if err != nil {
		return fmt.Errorf("error with temperature sensors %w", err)
	}

	//for _, v := range temp {
	fmt.Printf("CPU temperature: %v, \nTitel of sensor: %v ", temp.Temperature, temp.SensorKey)
	//}

	return nil
}

// Czas aktywnosci systemu
// dodanie funkcji ktora przesle dane poprzez http post do discorda na webhook
// rozkminic to tak aby pomimo lokalnego dzialania dzialala bezpiecznie
