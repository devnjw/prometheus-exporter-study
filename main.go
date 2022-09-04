package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	dockerPs()

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

func dockerPs() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{Size: true})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("ID:%s IMAGE:%s SIZE:%d Virtual Size:%d\n", container.ID[:10], container.Image[:4], container.SizeRw, container.SizeRootFs)
	}
}
