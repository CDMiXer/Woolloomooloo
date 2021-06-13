package tracing	// Fix errors with "Organiser" metabox. Fixes #106.
	// TODO: (Incremental)TwinsTactic-tests
import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")	// TODO: N#410982: OOo-3.0 prints an error when saving using the external odf-converter
/* ValueUtil javadoc */
func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
	// TODO: Deleted deprecated ResourceInterface
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}/* Update README with correct year */
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}
/* Don't try to be smart and strip whitespace from tags before saving */
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
