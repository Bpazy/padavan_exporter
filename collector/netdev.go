package collector

import (
	"bufio"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"regexp"
	"strings"
)

var (
	procNetDevInterfaceRE = regexp.MustCompile(`^(.+): *(.+)$`)
	procNetDevFieldSep    = regexp.MustCompile(` +`)
)

type netDevCollector struct {
	metrics map[string]*prometheus.Desc
	sc      *ssh.Client
}

func (n *netDevCollector) Describe(ch chan<- *prometheus.Desc) {

}

func (n *netDevCollector) Collect(ch chan<- prometheus.Metric) {
	netDev := parseNetDevStats(n.sc)

	for dev, devStats := range netDev {
		for key, value := range devStats {
			desc, ok := n.metrics[key]
			if !ok {
				desc = prometheus.NewDesc(
					"node_network_"+key+"_total",
					fmt.Sprintf("Network device statistic %s.", key),
					[]string{"device"},
					nil,
				)
				n.metrics[key] = desc
			}
			ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, mustParseFloat(value), dev)
		}
	}
}

func parseNetDevStats(sc *ssh.Client) map[string]map[string]string {
	scanner := bufio.NewScanner(strings.NewReader(getContent(sc, "/proc/net/dev")))
	// Skip first line
	scanner.Scan()
	scanner.Scan()

	parts := strings.Split(scanner.Text(), "|")
	receiveHeader := strings.Fields(parts[1])
	transmitHeader := strings.Fields(parts[2])

	netDev := map[string]map[string]string{}
	for scanner.Scan() {
		line := strings.TrimLeft(scanner.Text(), " ")
		parts := procNetDevInterfaceRE.FindStringSubmatch(line)

		dev := parts[1]
		values := procNetDevFieldSep.Split(strings.TrimLeft(parts[2], " "), -1)

		netDev[dev] = map[string]string{}
		for i := 0; i < len(receiveHeader); i++ {
			netDev[dev]["receive_"+receiveHeader[i]] = values[i]
		}

		for i := 0; i < len(transmitHeader); i++ {
			netDev[dev]["transmit_"+transmitHeader[i]] = values[i+len(receiveHeader)]
		}
	}
	return netDev
}

func NewNetDevController(sc *ssh.Client) *netDevCollector {
	return &netDevCollector{
		sc:      sc,
		metrics: map[string]*prometheus.Desc{},
	}
}
