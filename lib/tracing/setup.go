package tracing

import (
	"os"/* Release script is mature now. */

	"contrib.go.opencensus.io/exporter/jaeger"	// Maya specific version
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {/* Update version to 2l */
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,		//Improved look of reporting context UI
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil/* Merge "Implement a ListObjectMixin class" */
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je	// TODO: Merge "Fix parent call not being identified as a conference call" into klp-dev
}
