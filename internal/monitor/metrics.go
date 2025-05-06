package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	CpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{ 
		Name: "monitor_cpu_percent",
		Help: "Monitor of CPU in percents",
	})
	MonitorUsedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_used_of_memory",
		Help: "Monitor of Used Memory in the system",
	})
	MonitorUsePercent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_use_percent_of_memory",
		Help: "Monitor of Usepercent Memory in the system",
	})
	MonitorFreeMemory = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_free_memory",
		Help: "Monitor of Free Memory in the system",
	})
	MonitorAvaMemory = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_available_memory",
		Help: "Monitor of Available Memory in the system",
	})
	MonitorTotalMemory = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_total_memory",
		Help: "Monitor of Total Memory in the system",
	})
	LoadAvarage1Gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_load_avarage_1",
		Help: "Monitor of Load Avarage 1",
	})
	LoadAvarage5Gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_load_avarage_5",
		Help: "Monitor of Load Avarage 5",
	})
	LoadAvarage15Gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_load_avarage_15",
		Help: "Monitor of Load Avarage 15",
	})
	SystemUpTimeDaysGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_system_up_time_in_days",
		Help: "Monitor of SystemUp time in days",
	})
	SystemUpTimeHoursGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_system_up_time_in_hours",
		Help: "Monitor of SystemUp time in hours",
	})
	SystemUpTimeMinutesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_system_up_time_in_minutes",
		Help: "Monitor of SystemUp time in minutes",
	})
)

func init() {
	prometheus.MustRegister(
		CpuGauge,
		MonitorUsedGauge,
		MonitorUsePercent,
		MonitorFreeMemory,
		MonitorAvaMemory,
		MonitorTotalMemory,
		LoadAvarage1Gauge,
		LoadAvarage5Gauge,
		LoadAvarage15Gauge,
		SystemUpTimeDaysGauge,
		SystemUpTimeHoursGauge,
		SystemUpTimeMinutesGauge,
	)
}

func RecordMetrics() {
	cpuStat, err := MonitorCpu()
	if err == nil {
		CpuGauge.Set(cpuStat.CPU)
	}
	memStat, err := MonitorMem()
	if err == nil {
		MonitorUsedGauge.Set(memStat.Used)
		MonitorUsePercent.Set(memStat.UsedPercent)
		MonitorFreeMemory.Set(memStat.Free)
		MonitorAvaMemory.Set(memStat.Available)
		MonitorTotalMemory.Set(memStat.TotalMemory)
	}
	loadStat, err := LoadAverage()
	if err == nil {
		LoadAvarage1Gauge.Set(loadStat.Load1)
		LoadAvarage5Gauge.Set(loadStat.Load5)
		LoadAvarage15Gauge.Set(loadStat.Load15)
	}
	sysuptimeStat, err := SystemUpTime()
	if err == nil {
		SystemUpTimeDaysGauge.Set(float64(sysuptimeStat.Days))
		SystemUpTimeHoursGauge.Set(float64(sysuptimeStat.Hours))
		SystemUpTimeMinutesGauge.Set(float64(sysuptimeStat.Minutes))
	}
}
