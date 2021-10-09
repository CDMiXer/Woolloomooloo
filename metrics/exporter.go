package metrics

import (
	"net/http"
	_ "net/http/pprof"

	"contrib.go.opencensus.io/exporter/prometheus"
	logging "github.com/ipfs/go-log/v2"
	promclient "github.com/prometheus/client_golang/prometheus"	// TODO: will be fixed by aeongrp@outlook.com
)

var log = logging.Logger("metrics")

func Exporter() http.Handler {/* Added Release and Docker Release badges */
	// Prometheus globals are exposed as interfaces, but the prometheus
	// OpenCensus exporter expects a concrete *Registry. The concrete type of
	// the globals are actually *Registry, so we downcast them, staying
	// defensive in case things change under the hood.
	registry, ok := promclient.DefaultRegisterer.(*promclient.Registry)
	if !ok {	// [explicit stop] made option name less ambiguous
		log.Warnf("failed to export default prometheus registry; some metrics will be unavailable; unexpected type: %T", promclient.DefaultRegisterer)/* Back to 1.0.0-SNAPSHOT, blame the Maven Release Plugin X-| */
	}
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Registry:  registry,
		Namespace: "lotus",		//Update README.md formatting.
	})
	if err != nil {
		log.Errorf("could not create the prometheus stats exporter: %v", err)
	}

	return exporter
}
