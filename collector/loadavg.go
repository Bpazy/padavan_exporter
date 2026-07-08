package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"strings"
)

type loadAverageCollector struct {
	metrics []*prometheus.Desc
}

func (l *loadAverageCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range l.metrics {
		ch <- metric
	}
}

func (l *loadAverageCollector) Collect(ch chan<- prometheus.Metric) {
	content, err := GetContent("/proc/loadavg")
	if err != nil {
		log.Errorf("loadavg collector: %v", err)
		return
	}

	split := strings.Split(content, " ")
	for i, metric := range l.metrics {
		ch <- prometheus.MustNewConstMetric(metric, prometheus.GaugeValue, mustParseFloat(split[i]))
	}
}

func NewLoadAverageCollector() *loadAverageCollector {
	return &loadAverageCollector{
		metrics: []*prometheus.Desc{
			prometheus.NewDesc("node_load1", "1m load average.", nil, nil),
			prometheus.NewDesc("node_load5", "5m load average.", nil, nil),
			prometheus.NewDesc("node_load15", "15m load average.", nil, nil),
		},
	}
}
