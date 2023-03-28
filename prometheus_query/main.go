package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func main() {
	// Create a Prometheus API client
	client, err := api.NewClient(api.Config{
		Address: "http://127.0.0.1:9090",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Create a Prometheus query API client
	queryClient := v1.NewAPI(client)

	// Define the query
	query := "rate(promhttp_metric_handler_requests_total{code=\"200\"}[1m])"

	// Query Prometheus
	result, err := queryClient.Query(context.Background(), query, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range result.(model.Vector) {
		fmt.Printf("vector: %s %.2f\n", v.Metric, v.Value)
	}

	// Convert the result to a float64
	value := float64(result.(model.Vector)[0].Value)
	fmt.Printf("HTTP requests per second: %.2f\n", value)
}
