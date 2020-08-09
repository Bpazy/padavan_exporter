package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"padavan_exporter/collector"
	"time"
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
	reg := prometheus.NewPedanticRegistry()
	sc := initSshClient()
	reg.MustRegister(collector.NewLoadAverageCollector(sc))
	reg.MustRegister(collector.NewNetDevController(sc))

	gatherers := prometheus.Gatherers{reg}
	h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{
		ErrorLog:      log.StandardLogger(),
		ErrorHandling: promhttp.ContinueOnError,
	})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})

	log.Printf("Start server at %s", *la)
	log.Fatal(http.ListenAndServe(*la, nil))
}

func initSshClient() *ssh.Client {
	sshConfig := ssh.ClientConfig{
		User:            *pu,
		Auth:            []ssh.AuthMethod{ssh.Password(*pp)},
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	log.Printf("Connecting to %s", *ph)
	sshClient, err := ssh.Dial("tcp", *ph, &sshConfig)
	if err != nil {
		log.Fatalf("create ssh client failed: %+v", err)
	}
	return sshClient
}
