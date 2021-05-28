package tracing
/* + [cucmber] code cleaning */
import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")
/* test(suites): add link of benchmark suite */
func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
/* Release v1.1.0-beta1 (#758) */
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}/* Create doctrine.local.dist */
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,	// upgrade tcpdf to version: 6.0.055  - fonts
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}/* print jbig2dec warnings to stderr */

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})		//Improve std::string
	return je
}
