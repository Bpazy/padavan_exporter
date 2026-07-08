package collector

import (
	"bufio"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
)

var (
	meminfoReg = regexp.MustCompile(`(\w+):\s+(\d+) kB`)
)

type memoryCollector struct {
	metrics map[string]*prometheus.Desc
}

func (m *memoryCollector) Describe(ch chan<- *prometheus.Desc) {
	// metrics created when Collect
}

func (m *memoryCollector) Collect(ch chan<- prometheus.Metric) {
	content, err := GetContent("/proc/meminfo")
	if err != nil {
		log.Errorf("memory collector: %v", err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		parts := meminfoReg.FindStringSubmatch(scanner.Text())

		if len(parts) != 3 {
			continue
		}
		key := parts[1]
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
		}
		value *= 1024 // convert kB to B

		var desc *prometheus.Desc
		var ok bool
		switch key {
		case "MemTotal":
			desc, ok = m.metrics["MemTotal"]
			if !ok {
				desc = prometheus.NewDesc("node_memory_total_bytes", "Total memory in bytes.", nil, nil)
				m.metrics["MemTotal"] = desc
			}
		case "MemFree":
			desc, ok = m.metrics["MemFree"]
			if !ok {
				desc = prometheus.NewDesc("node_memory_free_bytes", "Free memory in bytes.", nil, nil)
				m.metrics["MemFree"] = desc
			}
		case "Buffers":
			desc, ok = m.metrics["Buffers"]
			if !ok {
				desc = prometheus.NewDesc("node_memory_buffers_bytes", "Buffers memory in bytes.", nil, nil)
				m.metrics["Buffers"] = desc
			}
		case "Cached":
			desc, ok = m.metrics["Cached"]
			if !ok {
				desc = prometheus.NewDesc("node_memory_cached_bytes", "Cached memory in bytes.", nil, nil)
				m.metrics["Cached"] = desc
			}
		default:
			continue
		}
		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, value)
	}
}

func NewMemoryCollector() *memoryCollector {
	return &memoryCollector{
		metrics: map[string]*prometheus.Desc{},
	}
}
