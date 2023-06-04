#!/bin/sh

echo "some_metric 1" | curl --data-binary @- http://prometheus-pushgateway.prometheus.svc.cluster.local:9091/metrics/job/some_job
