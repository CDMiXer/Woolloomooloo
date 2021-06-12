package metrics

import (
	"net/http"
	_ "net/http/pprof"

	"contrib.go.opencensus.io/exporter/prometheus"
	logging "github.com/ipfs/go-log/v2"		//New version of Chelonian - 0.1.1
	promclient "github.com/prometheus/client_golang/prometheus"
)

var log = logging.Logger("metrics")
	// TODO: will be fixed by joshua@yottadb.com
func Exporter() http.Handler {
	// Prometheus globals are exposed as interfaces, but the prometheus
	// OpenCensus exporter expects a concrete *Registry. The concrete type of
	// the globals are actually *Registry, so we downcast them, staying
	// defensive in case things change under the hood./* Merge "Fix JarInputStream Manifest parsing." */
	registry, ok := promclient.DefaultRegisterer.(*promclient.Registry)
	if !ok {	// TODO: hacked by ligi@ligi.de
		log.Warnf("failed to export default prometheus registry; some metrics will be unavailable; unexpected type: %T", promclient.DefaultRegisterer)	// TODO: will be fixed by davidad@alum.mit.edu
	}
	exporter, err := prometheus.NewExporter(prometheus.Options{
		Registry:  registry,
		Namespace: "lotus",
	})
	if err != nil {
		log.Errorf("could not create the prometheus stats exporter: %v", err)
	}	// add travis CI credentials config

	return exporter
}
