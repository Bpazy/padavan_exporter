package collector

import (
	log "github.com/sirupsen/logrus"
	"strconv"
)

func mustParseFloat(fs string) float64 {
	float, err := strconv.ParseFloat(fs, 32)
	if err != nil {
		log.Printf("%+v", err)
		return 0
	}
	return float
}
