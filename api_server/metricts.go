package api_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/conexxxion/conexxxion-backoffice/config"
)

func StartMetricsServer() {
	port := config.GetConfig().MetricsPort
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("starting metrics server: ", err)
	}
}
