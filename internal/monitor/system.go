package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v4/sensors"
)

// !! wszystkie fmt.Println do wywalnie na koncu

// dodac jeszcze inne ciekawe funkcji ktore moga byc atrakcyjne w typ wypadku
type DataStat struct { // dodanie tego zeby wyprowadzilo metryki z dockera prometheusa i dodalo do grafany
	CPU         float64 `json:"cpu"`
	Memory      float64 `json:"memory"` //tutaj mozna zostawic to jako unit64
	Available   float64 `json:"available"`
	Used        float64 `json:"used"`
	Free        float64 `json:"free"`
	UsedPercent float64 `json:"usedpercent"`
	Load1       float64 `json:"load1"`
	Load5       float64 `json:"load5"`
	Load15      float64 `json:"load15"`
	CpuTemp     float64 `json:"cputemp"`
	Days        int     `json:"days"`
	Hours       int     `json:"hours"`
	Minutes     int     `json:"minutes"`

	//Temperature float64 `json:"temperature"` // możesz dodać więcej jeśli masz wiele sensorów
	//Uptime      string  `json:"uptime"`
	//Timestamp   string  `json:"timestamp"`
}

// CPU
func MonitorCpu() (*DataStat, error) {

	usage, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	data := &DataStat{
		CPU: usage[0],
	}

	if data.CPU > 80.0 {
		fmt.Println("--Too high CPU!--") // dodac to do webhooka
	}
	return data, nil
}

// Memory
func MonitorMem() (*DataStat, error) {

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	used := float64(v.Used)
	usePerc := float64(v.UsedPercent)
	freeMemory := float64(v.Free)
	avaMemory := float64(v.Available)
	totalMemory := float64(v.Total)

	data := &DataStat{
		Available:   avaMemory,
		Used:        used,
		Memory:      totalMemory,
		UsedPercent: usePerc,
		Free:        freeMemory,
	}
	return data, nil
}

// Load Average obciazenie systemu
func LoadAverage() (*DataStat, error) {

	lavg, err := load.Avg()
	if err != nil {
		return nil, err
	}

	data := &DataStat{
		Load1:  lavg.Load1,
		Load5:  lavg.Load5,
		Load15: lavg.Load15,
	}
	fmt.Printf("Load Average: 1m: %.2f, \nLoad Average 5m: %.2f \nLoad Average 15m: %.2f", lavg.Load1, lavg.Load5, lavg.Load15)
	return data, nil
}

// CPU temperature
func TempCpu() (*DataStat, error) {

	var cputemp float64

	temp, err := sensors.SensorsTemperatures()
	if err != nil {
		return nil, err
	}

	for _, v := range temp {
		fmt.Printf("CPU temperature: %v, \nTitel of sensor: %v ", v.Temperature, v.SensorKey)
	}

	//cputemp = v.Temperature

	data := &DataStat{
		CpuTemp: cputemp,
	}
	return data, nil
}

// Czas aktywnosci systemu
func SystemUpTime() (*DataStat, error) {

	uptime, err := host.Uptime()
	if err != nil {
		return nil, err
	}
	days := int(uptime) / 86400
	hours := (int(uptime) % 86400) / 3600
	minutes := (int(uptime) % 3600) / 60

	dataCpuTemp := &DataStat{
		Days:    days, // tu skonczylem, jakis blad ?
		Hours:   hours,
		Minutes: minutes,
	}
	return dataCpuTemp, nil
}

// dodanie funkcji ktora przesle dane poprzez http post do discorda na webhook // to jako server.go bedzie
// rozkminic to tak aby pomimo lokalnego dzialania dzialala bezpiecznie
