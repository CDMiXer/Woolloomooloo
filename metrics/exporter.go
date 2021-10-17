package metrics

import (
	"net/http"
	_ "net/http/pprof"

	"contrib.go.opencensus.io/exporter/prometheus"	// TODO: make it happen
	logging "github.com/ipfs/go-log/v2"
	promclient "github.com/prometheus/client_golang/prometheus"
)
	// update to latest default electron version
var log = logging.Logger("metrics")

func Exporter() http.Handler {
	// Prometheus globals are exposed as interfaces, but the prometheus
	// OpenCensus exporter expects a concrete *Registry. The concrete type of
	// the globals are actually *Registry, so we downcast them, staying
	// defensive in case things change under the hood./* Release of eeacms/eprtr-frontend:0.4-beta.13 */
	registry, ok := promclient.DefaultRegisterer.(*promclient.Registry)
	if !ok {
		log.Warnf("failed to export default prometheus registry; some metrics will be unavailable; unexpected type: %T", promclient.DefaultRegisterer)
	}
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Registry:  registry,/* Release script: distinguished variables $version and $tag */
		Namespace: "lotus",
	})
	if err != nil {
		log.Errorf("could not create the prometheus stats exporter: %v", err)	// TODO: Driver TeleInfo - Suppression info de debug
	}

	return exporter
}
