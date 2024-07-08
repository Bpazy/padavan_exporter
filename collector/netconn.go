package collector

import (
	"bufio"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"gopkg.in/alecthomas/kingpin.v2"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  netconn.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-07-08 10:10
 */
const (
	namespace         = "node"
	netStatsSubsystem = "netstat"
)

var (
	netStatFields = kingpin.Flag(
		"collector.netstat.fields",
		"Regexp of fields to return for netstat collector.",
	).Default(
		"^(.*_(InErrors|InErrs)|Ip_Forwarding|Ip(6|Ext)_(InOctets|OutOctets)|Icmp6?_(InMsgs|OutMsgs)|TcpExt_(Listen.*|Syncookies.*|TCPSynRetrans|TCPTimeouts|TCPOFOQueue)|Tcp_(ActiveOpens|InSegs|OutSegs|OutRsts|PassiveOpens|RetransSegs|CurrEstab)|Udp6?_(InDatagrams|OutDatagrams|NoPorts|RcvbufErrors|SndbufErrors))$",
	).String()
)

type NetconnCollector struct {
	metrics      map[string]*prometheus.Desc // Stores metric descriptions
	sc           *ssh.Client                 // SSH client for remote data collection
	fieldPattern *regexp.Regexp              // Regex for matching field names
}

func (c *NetconnCollector) Describe(ch chan<- *prometheus.Desc) {
	// Metrics are created during Collect
}

func (c *NetconnCollector) Collect(ch chan<- prometheus.Metric) {
	netStats := parseNetStats(mustGetContent(c.sc, "/proc/net/netstat"))
	snmpStats := parseNetStats(mustGetContent(c.sc, "/proc/net/snmp"))
	snmp6Stats := parseSNMP6Stats(mustGetContent(c.sc, "/proc/net/snmp6"))

	// Merge snmpStats and snmp6Stats into netStats
	for k, v := range snmpStats {
		netStats[k] = v
	}
	for k, v := range snmp6Stats {
		netStats[k] = v
	}

	for protocol, protocolStats := range netStats {
		for name, value := range protocolStats {
			key := protocol + "_" + name
			if !c.fieldPattern.MatchString(key) {
				continue
			}

			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				fmt.Printf("Invalid value in netstats: %v, error: %v\n", value, err)
				continue
			}

			desc := prometheus.NewDesc(
				prometheus.BuildFQName(namespace, netStatsSubsystem, key),
				fmt.Sprintf("Statistic %s.", key),
				nil, nil,
			)

			ch <- prometheus.MustNewConstMetric(desc, prometheus.UntypedValue, v)
		}
	}
}

func parseNetStats(content string) map[string]map[string]string {
	// 将字符串内容转换为 io.Reader
	reader := strings.NewReader(content)
	netStats := make(map[string]map[string]string)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		nameParts := strings.Split(scanner.Text(), " ")
		if !scanner.Scan() {
			break
		}
		valueParts := strings.Split(scanner.Text(), " ")
		protocol := strings.TrimSuffix(nameParts[0], ":")
		netStats[protocol] = make(map[string]string)
		if len(nameParts) != len(valueParts) {
			continue
		}
		for i := 1; i < len(nameParts); i++ {
			netStats[protocol][nameParts[i]] = valueParts[i]
		}
	}

	return netStats
}

func parseSNMP6Stats(content string) map[string]map[string]string {
	reader := strings.NewReader(content)

	netStats := make(map[string]map[string]string)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		stat := strings.Fields(scanner.Text())
		if len(stat) < 2 {
			continue
		}
		if sixIndex := strings.Index(stat[0], "6"); sixIndex != -1 {
			protocol := stat[0][:sixIndex+1]
			name := stat[0][sixIndex+1:]
			if _, present := netStats[protocol]; !present {
				netStats[protocol] = make(map[string]string)
			}
			netStats[protocol][name] = stat[1]
		}
	}

	return netStats
}

func NewNetconnCollector(sc *ssh.Client) *NetconnCollector {
	pattern := regexp.MustCompile(*netStatFields)
	return &NetconnCollector{
		sc:           sc,
		metrics:      map[string]*prometheus.Desc{},
		fieldPattern: pattern,
	}
}
