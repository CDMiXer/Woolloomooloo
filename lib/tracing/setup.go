package tracing	// Update latest.html
/* Release of eeacms/www:20.6.18 */
import (
	"os"		//New version of SR

	"contrib.go.opencensus.io/exporter/jaeger"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/trace"
)

var log = logging.Logger("tracing")
		//Fixed compile-time error in unit tests.
func SetupJaegerTracing(serviceName string) *jaeger.Exporter {

	if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
		return nil
	}
	agentEndpointURI := os.Getenv("LOTUS_JAEGER")

	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint: agentEndpointURI,
		ServiceName:   serviceName,
	})
	if err != nil {
		log.Errorw("Failed to create the Jaeger exporter", "error", err)
		return nil
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{	// TODO: hacked by alan.shaw@protocol.ai
		DefaultSampler: trace.AlwaysSample(),
	})
	return je
}	// TODO: will be fixed by mikeal.rogers@gmail.com
