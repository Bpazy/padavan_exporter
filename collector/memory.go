package collector

import (
	"bufio"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/crypto/ssh"
	"regexp"
	"strconv"
	"strings"
)

/**
 * @Author: 南宫乘风
 * @Description:
 * @File:  memory.go
 * @Email: 1794748404@qq.com
 * @Date: 2024-05-22 15:05
 */

var (
	meminfoReg = regexp.MustCompile(`(\w+):\s+(\d+) kB`)
)

// memoryCollector 收集内存相关指标信息
type memoryCollector struct {
	metrics map[string]*prometheus.Desc // 存储指标描述信息
	sc      *ssh.Client                 // SSH客户端，用于远程收集指标信息
}

// Describe 向通道发送描述指标的Desc对象
// 这个方法不具体实现，因为metrics在Collect方法中动态创建。
func (m *memoryCollector) Describe(ch chan<- *prometheus.Desc) {
	// metrics created when Collect
}

// Collect 向通道发送内存指标数据
// @param ch 用于发送指标的通道
func (m *memoryCollector) Collect(ch chan<- prometheus.Metric) {
	// 通过SSH客户端获取远程机器的/proc/meminfo文件内容
	content := mustGetContent(m.sc, "/proc/meminfo")
	// 创建Scanner用于逐行读取内容
	scanner := bufio.NewScanner(strings.NewReader(content))
	// 使用正则表达式匹配行内容，获取指标名和值
	for scanner.Scan() {
		parts := meminfoReg.FindStringSubmatch(scanner.Text())

		if len(parts) != 3 {
			// 如果匹配不成功，则跳过当前行
			continue
		}
		// 指标名
		key := parts[1]
		// 解析指标值
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			continue
			// 如果解析失败，则跳过当前行
		}
		value *= 1024 // 将值从kB转换为B

		var desc *prometheus.Desc
		var ok bool
		// 根据指标名，获取或创建对应的Desc对象
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
		case "MemAvailable":
			desc, ok = m.metrics["MemAvailable"]
			if !ok {
				desc = prometheus.NewDesc("node_memory_available_bytes", "Available memory in bytes.", nil, nil)
				m.metrics["MemAvailable"] = desc
			}
		default:
			continue // 如果不是我们关心的指标，则跳过当前行
		}
		// 向通道发送指标
		ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, value)
	}
}

// NewMemoryCollector 创建一个新的memoryCollector实例
// @param sc SSH客户端
// @return 返回memoryCollector实例的指针
func NewMemoryCollector(sc *ssh.Client) *memoryCollector {
	return &memoryCollector{
		sc:      sc,
		metrics: map[string]*prometheus.Desc{},
	}
}
