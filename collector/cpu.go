package collector

import (
	"bufio"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

var (
	procStatReg = regexp.MustCompile("(cpu.+?) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+)")
)

type cpuCollector struct {
	metrics map[string]*prometheus.Desc
}

func (s *cpuCollector) Describe(ch chan<- *prometheus.Desc) {
	// metrics created when Collect
}

func (s *cpuCollector) Collect(ch chan<- prometheus.Metric) {
	content, err := GetContent("/proc/stat")
	if err != nil {
		log.Errorf("cpu collector: %v", err)
		return
	}

	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		parts := procStatReg.FindStringSubmatch(scanner.Text())
		if len(parts) != 12 {
			continue
		}
		dev := strings.TrimSpace(parts[1])
		user := parts[2]
		system := parts[3]
		idle := parts[4]
		iowait := parts[5]
		irq := parts[6]
		softirq := parts[7]

		cachedDesc, ok := s.metrics[dev]
		if !ok {
			cachedDesc = prometheus.NewDesc(
				"node_cpu_seconds_total",
				"Seconds the cpus spent in each mode.",
				[]string{"cpu", "mode"},
				nil,
			)
			s.metrics[dev] = cachedDesc
		}
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(user), dev, "user")
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(system), dev, "system")
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(idle), dev, "idle")
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(iowait), dev, "iowait")
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(irq), dev, "irq")
		ch <- prometheus.MustNewConstMetric(cachedDesc, prometheus.CounterValue, mustParseFloat(softirq), dev, "softirq")
	}
}

func NewCpuCollector() *cpuCollector {
	return &cpuCollector{
		metrics: map[string]*prometheus.Desc{},
	}
}
