package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {

	sample := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "my_custom_metric",
		Help: "my custom metrics",
	})
	sample.Inc()
	if err := push.New("http://localhost:9091", "my_job").
		Collector(sample).
		Push(); err != nil {
		fmt.Printf(err.Error())
	}
}
