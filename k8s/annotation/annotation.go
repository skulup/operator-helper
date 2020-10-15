package annotation

import (
	"strconv"
)

const (
	annPrometheusScrape = "prometheus.io/scrape"
	annPrometheusPort   = "prometheus.io/port"
)

// DecorateForPrometheus adds prometheus scraping annotations
func DecorateForPrometheus(ann map[string]string, scrap bool, port int) map[string]string {
	ann[annPrometheusScrape] = strconv.FormatBool(scrap)
	if scrap {
		ann[annPrometheusPort] = strconv.FormatInt(int64(port), 10)
	}
	return ann
}
