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

func init() {
	la = kingpin.Flag("web.listen-address", "Address on which to expose metrics and web interface").Default(":9100").String()
	ph = kingpin.Flag("padavan.ssh.host", "Padavan ssh host").Default("127.0.0.1:22").String()
	pu = kingpin.Flag("padavan.ssh.username", "Padavan ssh username").Default("admin").String()
	pp = kingpin.Flag("padavan.ssh.password", "Padavan ssh password").Default("admin").String()
	isDebug := kingpin.Flag("debug", "Debug mode").Bool()
	kingpin.Parse()

	if *isDebug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugf("web.listen-address(%s) padavan.ssh.host(%s) padavan.ssh.username(%s) padavan.ssh.password(%s)", *la, *ph, *pu, *pp)
}

func main() {
	collector.SetSSHConfig(*ph, *pu, *pp)

	// Wait for initial SSH connection with automatic retry.
	log.Infof("Connecting to %s ...", *ph)
	collector.WaitForConnection()

	reg := prometheus.NewPedanticRegistry()
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
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, homePage())
	})

	log.Printf("Start server at %s", *la)
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
