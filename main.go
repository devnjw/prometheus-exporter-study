package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// custom registry to get rid of additional go_collector metrics
	r := prometheus.NewRegistry()
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	labels := make(map[string]string)
	labels["instance_id"] = "265"

	metricName := "INSTANCE_SIZE"
	metric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        metricName,
			ConstLabels: labels,
			Help:        "Custom metric",
		},
	)
	r.MustRegister(metric)
	metric.Set(float64(0))

	http.Handle("/metrics", handler)
	log.Printf("Starting to listen on :%d", 9090)
	err := http.ListenAndServe(fmt.Sprintf(":%d", 9090), nil)
	if err != nil {
		log.Println("Failed to Listen And Serve")
	}
}
