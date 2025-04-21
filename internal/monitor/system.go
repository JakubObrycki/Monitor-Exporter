package monitor

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v4/sensors"
)

// !! wszystkie fmt.Println do wywalnie na koncu

// dodac jeszcze inne ciekawe funkcji ktore moga byc atrakcyjne w typ wypadku
type DataStat struct { // dodanie tego zeby wyprowadzilo metryki z dockera prometheusa i dodalo do grafany
	CPU         float64 `json:"cpu"`
	Memory      float64 `json:"memory"`
	Available   float64 `json:"available"`
	Used        float64 `json: used`
	Free        float64 `json:"free"`
	UsedPercent float64 `json:"usedpercent"`
	//Load1       float64 `json:"load1"`
	//Load5       float64 `json:"load5"`
	//Load15      float64 `json:"load15"`
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

	roundvalue := math.Round(usage[0]*100) / 100

	data := &DataStat{
		CPU: roundvalue,
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

	used := float64(v.Used)           // w sumie to tez bedzie do wywalenie bo do prometheusa trzeba dddawac w bajtach
	usePerc := float64(v.UsedPercent) // tu bedzie alert, podbnie tez dla grafany (pow.80%)
	freeMemory := float64(v.Free)     // skonczylem tu free memorys < 1GB
	avaMemory := float64(v.Available) // tu bedzie alert ale najwyzej to dla grafany sie doda
	totalMemory := float64(v.Total)

	data := &DataStat{
		Available:   avaMemory,
		Used:        used,
		Memory:      totalMemory,
		UsedPercent: usePerc,
		Free:        freeMemory,
	}

	fmt.Println("Free memory: ", freeMemory)
	fmt.Println("Used: ", used, "GB")
	fmt.Println("Use percent: ", usePerc, "%")
	fmt.Println("Available memory:", avaMemory, "GB")
	fmt.Printf("Total Memory: %.2f GB\n", data.Memory) // oproc total memory, mozna dodac cos jeszcze
	if data.Memory > 80.0 {                            // alert musi byc ustawiony do usedpercent !
		fmt.Println("--Too high Memory!--") // dodac to do webhooka, trzeba napisac dodatkowa funkcje do webhookow
	}

	return data, nil
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

	temp, err := sensors.SensorsTemperatures()
	if err != nil {
		return fmt.Errorf("error with temperature sensors %w", err)
	}

	for _, v := range temp {
		fmt.Printf("CPU temperature: %v, \nTitel of sensor: %v ", v.Temperature, v.SensorKey)
	}

	return nil
}

// Czas aktywnosci systemu
// dodanie funkcji ktora przesle dane poprzez http post do discorda na webhook
// rozkminic to tak aby pomimo lokalnego dzialania dzialala bezpiecznie
