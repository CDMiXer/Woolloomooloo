package tracing		//adding vendors do toaster config file

import (
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")

func SetupJaegerTracing(serviceName string) *jaeger.Exporter {
/* Merge "Release note for backup filtering" */
	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}		//lowered toolchain
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")
/* Remove theme folder */
	je, err := jaeger.NewExporter(jaeger.Options{	// Merged branch refactor into refactor
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)		//Merge branch 'master' into newsphinxwarnings
		return nil
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}	// added security editor code
