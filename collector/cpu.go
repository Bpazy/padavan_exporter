package collector

import (
	"bufio"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"regexp"
	"strings"
)

var (
	procStatReg = regexp.MustCompile("(cpu.+?)\\s+(\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+) (\\d+)")
)

type cpuCollector struct {
	metrics map[string]*prometheus.Desc
	sc      *ssh.Client
}

func (s *cpuCollector) Describe(ch chan<- *prometheus.Desc) {
	// metrics created when Collect
}

func (s *cpuCollector) Collect(ch chan<- prometheus.Metric) {
	scanner := bufio.NewScanner(strings.NewReader(getContent(s.sc, "/proc/stat")))
	scanner.Scan()

	for scanner.Scan() {
		parts := procStatReg.FindStringSubmatch(scanner.Text())
		if len(parts) != 12 {
			continue
		}
		dev := parts[1]
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

func NewCpuCollector(sc *ssh.Client) *cpuCollector {
	return &cpuCollector{
		sc:      sc,
		metrics: map[string]*prometheus.Desc{},
	}
}
