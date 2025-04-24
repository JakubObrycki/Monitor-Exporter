package monitor

import (
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
)

type DataStat struct {
	CPU         float64 `json:"cpu"`
	TotalMemory float64 `json:"memory"`
	Available   float64 `json:"available"`
	Used        float64 `json:"used"`
	Free        float64 `json:"free"`
	UsedPercent float64 `json:"usedpercent"`
	Load1       float64 `json:"load1"`
	Load5       float64 `json:"load5"`
	Load15      float64 `json:"load15"`
	Days        int     `json:"days"`
	Hours       int     `json:"hours"`
	Minutes     int     `json:"minutes"`
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

	//if data.CPU > 80.0 {
	//	fmt.Println("--Too high CPU!--") // dodac to do webhooka
	//}
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
		TotalMemory: totalMemory,
		UsedPercent: usePerc,
		Free:        freeMemory,
	}
	return data, nil
}

// Load Average
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
	return data, nil
}

// System activity time
func SystemUpTime() (*DataStat, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return nil, err
	}
	days := int(uptime) / 86400
	hours := (int(uptime) % 86400) / 3600
	minutes := (int(uptime) % 3600) / 60

	data := &DataStat{
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
	}
	return data, nil
}
