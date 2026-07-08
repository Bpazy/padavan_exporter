package main

import (
	"fmt"
	"github.com/Bpazy/padavan_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

var (
	ph *string // Padavan address
	pu *string // Padavan username
	pp *string // Padavan password
	la *string // Address on which to expose metrics and web interface
)

// init 函数用于初始化命令行参数。
// 通过 kingpin 解析命令行参数，包括 web 服务监听地址、Padavan SSH 主机地址、用户名、密码等。
// 并根据是否开启 debug 模式设置日志级别。
func init() {
	// 定义命令行参数
	la = kingpin.Flag("web.listen-address", "Address on which to expose metrics and web interface").Default(":9100").String()
	ph = kingpin.Flag("padavan.ssh.host", "Padavan ssh host").Default("127.0.0.1:22").String()
	pu = kingpin.Flag("padavan.ssh.username", "Padavan ssh username").Default("admin").String()
	pp = kingpin.Flag("padavan.ssh.password", "Padavan ssh password").Default("admin").String()
	// 解析命令行参数，并在 debug 模式下设置日志级别
	isDebug := kingpin.Flag("debug", "Debug mode").Bool()
	kingpin.Parse()

	if *isDebug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("web.listen-address(%s) padavan.ssh.host(%s) padavan.ssh.username(%s) padavan.ssh.password(%s)", *la, *ph, *pu, *pp)
}

// main 函数作为程序入口点，
// 负责建立 SSH 连接（带自动重试）、初始化 Prometheus 监控数据收集器，并启动 HTTP 服务提供监控数据。
func main() {
	// 配置 SSH 连接参数，并等待首次连接成功（失败时每 5 秒重试）
	collector.SetSSHConfig(*ph, *pu, *pp)
	log.Infof("Connecting to %s ...", *ph)
	collector.WaitForConnection()

	// 初始化 Prometheus 注册表
	reg := prometheus.NewPedanticRegistry()
	// 注册各种收集器到 Prometheus（收集器内部自动管理 SSH 连接）
	reg.MustRegister(collector.NewLoadAverageCollector())
	reg.MustRegister(collector.NewNetDevController())
	reg.MustRegister(collector.NewCpuCollector())
	reg.MustRegister(collector.NewMemoryCollector())
	reg.MustRegister(collector.NewNetconnCollector())

	gatherers := prometheus.Gatherers{reg}
	h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{
		ErrorLog:      log.StandardLogger(),
		ErrorHandling: promhttp.ContinueOnError,
	})
	// 处理 /metrics 请求以提供监控数据
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	// 处理根路径请求，提供简单信息页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, homePage())
	})

	log.Printf("Start server at %s", *la)
	// 启动 HTTP 服务
	log.Fatal(http.ListenAndServe(*la, nil))
}

func homePage() string {
	return `
<html>
 <head>
  <title>padavan_exporter</title>
 </head>
 <body>
  <h2>padavan_exporter</h2>
  <span>See docs at <a href="https://github.com/Bpazy/padavan_exporter">https://github.com/Bpazy/padavan_exporter</a></span>
  <br>
  <br>
  <span> Useful endpoints: </span>
  <br>
  <span> <a href="/metrics">metrics</a> <span> - available service metrics </span>
 </body>
</html>`
}
