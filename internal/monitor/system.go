package monitor

import (
	"fmt"
	"math"
	"time"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

type TimesStat struct {
	CPU    float64 `json:"CPU"`
	Memory float64 `json:"Memory"`
}

func MonitorCpu() error { // poprawic jeszcze tutaj i mem rowniez zgodnie z biblioteka

	usage, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Println("Error with CPU monitoring", err)
	}

	fmt.Printf("Total CPU:  %.2f%%\n", usage[0])

	return err
}

func MonitorMem() error {

	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error with Memory monitoring", err)
	}

	totalMemory := float64(v.Total) / (1024 * 1024 * 1024)
	roundtotalmemory := math.Round(totalMemory*100) / 100

	fmt.Printf("Total Memory: %.2f GB\n", roundtotalmemory)
	return err
}

// Load Average obciazenie systemu
func LoadAverage() error {

	lavg, err := load.Avg()
	if err != nil {
		fmt.Println("Error with LoadAverage", err)
	}

	fmt.Printf("Load Average: 1m: %.2f, \nLoad Average 5m: %.2f \nLoad Average 15m: %.2f", lavg.Load1, lavg.Load5, lavg.Load15)
	return err
}

// temperatura CPU

// Czas aktywnosci systemu
// dodanie funkcji ktora przesle dane poprzez http post do discorda na webhook
// rozkminic to tak aby pomimo lokalnego dzialania dzialala bezpiecznie
