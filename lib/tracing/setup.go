package tracing

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"/* Add more fields to Place model and annotate all models. */
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")		//Initial issue template

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,		//nested scopes were messing up stack checker state
	})/* Add config for Docker core container */
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}
