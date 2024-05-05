// main.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	memUtilization = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "memory_utilization",
			Help: "Current memory utilization of the host",
		},
	)
	cpuUtilization = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_utilization",
			Help: "Current CPU utilization of the host",
		},
	)
)

func init() {
	// Register the metrics with Prometheus
	prometheus.MustRegister(memUtilization)
	prometheus.MustRegister(cpuUtilization)
}

func updateMetrics() {
	// Update memory utilization metric
	vmStat, _ := mem.VirtualMemory()
	memUtilization.Set(vmStat.UsedPercent)

	// Update CPU utilization metric
	cpuStat, _ := cpu.Percent(0, false)
	cpuUtilization.Set(cpuStat[0])
}

func main() {
	// Start updating metrics in the background
	go func() {
		for {
			updateMetrics()
			time.Sleep(5 * time.Second) // Update metrics every 5 seconds
		}
	}()

	// Expose the registered metrics via HTTP
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
