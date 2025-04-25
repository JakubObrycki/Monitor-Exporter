package monitor

import (
	"github.com/prometheus/client_golang/prometheus" // blad dla wersji prometheus, pobawic sie z poprawna wersja Go
)

var (
	CpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{ // CpuGauge ma dostawac informacje z funkcja, dlatego funkcje nalezy zmienic lekko tak jak GPT zaproponwoal ale moim rozwiazaniem jest struct czyli niech pobiera z struct
		Name: "monitor_cpu_percent",        // nazwa metryki
		Help: "Monitor of CPU in percents", // opis metryki a dokladnie co metryka mierzy
	})

	MonitorUsedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_used_of_memory",               // nazwa metryki
		Help: "Monitor of Used Memory in the system", // opis metryki a dokladnie co metryka mierzy
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
	// tu skonczyelem, poprawic ponizsza metryke i dodac kolejne dotyczace tej funkcji + dla systemupTiem i tu chyba uzyc counter
	LoadAvarageGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "monitor_cpu_percent",        // nazwa metryki
		Help: "Monitor of CPU in percents", // opis metryki a dokladnie co metryka mierzy
	})
)
