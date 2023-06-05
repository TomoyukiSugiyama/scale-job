package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {

	sample := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "some_metric",
		Help: "my custom metrics",
	})
	sample.Add(3)
	if err := push.New("http://localhost:9091", "some_job").
		Collector(sample).
		Push(); err != nil {
		fmt.Printf(err.Error())
	}
}
