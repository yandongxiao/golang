package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	myGague = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        "my_example_gauge_data",
		Help:        "my example gauge data",
		ConstLabels: map[string]string{"error": ""},
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			myGague.Add(11)
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())

	recordMetrics()
	log.Fatal(http.ListenAndServe(*addr, nil))
}
