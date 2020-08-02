package main

import (
	"github.com/go-resty/resty/v2"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	load1 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "padavan_load1",
		Help: "Padavan 1 min load",
	})
	load5 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "padavan_load5",
		Help: "Padavan 5 min load",
	})
	load15 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "padavan_load15",
		Help: "Padavan 15 min load",
	})
	ph *string // Padavan address
	pu *string // Padavan username
	pp *string // Padavan password
	la *string // Address on which to expose metrics and web interface
)

var (
	client = resty.New()
)

func init() {
	la = kingpin.Flag("web.listen-address", "Address on which to expose metrics and web interface").Default(":9100").String()
	ph = kingpin.Flag("padavan.host", "Padavan address").Default("http://192.168.31.1").String()
	pu = kingpin.Flag("padavan.username", "Padavan username").Default("admin").String()
	pp = kingpin.Flag("padavan.password", "Padavan password").Default("admin").String()
	kingpin.Parse()

	client.SetBasicAuth(*pu, *pp)
	client.SetDisableWarn(true)
}

func main() {
	checkPadavan()
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(*la, nil)
	if err != nil {
		panic(err)
	}
}

// 校验 Padavan 地址、认证
func checkPadavan() {
	res, err := client.R().Get(getSystemStatusDataUrl())
	if err != nil {
		log.Fatalf("Connecting padava failed. Please check \"padavan.host\".")
	}
	if strings.Contains(res.String(), "401") {
		log.Fatalf("Authenticate failed. Please check \"padavan.username\" and \"padavan.password\".")
	}
}

func recordMetrics() {
	lReg := regexp.MustCompile("(\\d+\\.\\d+) (\\d+\\.\\d+) (\\d+\\.\\d+)")
	go func() {
		for {
			res, err := client.R().Get(getSystemStatusDataUrl())
			if err != nil {
				log.Printf("request system_status_data failed: %+v\n", err)
				time.Sleep(2 * time.Second)
				continue
			}

			loadStr := lReg.FindStringSubmatch(res.String())
			load1.Set(mustParseFloat(loadStr[1]))
			load5.Set(mustParseFloat(loadStr[2]))
			load15.Set(mustParseFloat(loadStr[3]))
			time.Sleep(2 * time.Second)
		}
	}()
}

func getSystemStatusDataUrl() string {
	return *ph + "/system_status_data.asp"
}

func mustParseFloat(fs string) float64 {
	float, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		log.Printf("%+v", err)
		return 0
	}
	return float
}
