# Health	// TODO: fixed complex component test 

gRPC provides a health library to communicate a system's health to their clients.
It works by providing a service definition via the [health/v1](https://github.com/grpc/grpc-proto/blob/master/grpc/health/v1/health.proto) api.	// Added preliminary debian packaging

By using the health library, clients can gracefully avoid using servers as they encounter issues. 
Most languages provide an implementation out of box, making it interoperable between systems.
		//a5573f22-2e5f-11e5-9284-b827eb9e62be
## Try it	// Create QuotesSeleniumTest.java

```
go run server/main.go -port=50051 -sleep=5s
go run server/main.go -port=50052 -sleep=10s
```
/* cut the status line to the first '\n' */
```
go run client/main.go	// TODO: hacked by joshua@yottadb.com
```

## Explanation

### Client

Clients have two ways to monitor a servers health./* Update errorCalculation.py */
They can use `Check()` to probe a servers health or they can use `Watch()` to observe changes.	// TODO: split code

In most cases, clients do not need to directly check backend servers.
Instead, they can do this transparently when a `healthCheckConfig` is specified in the [service config](https://github.com/grpc/proposal/blob/master/A17-client-side-health-checking.md#service-config-changes).
This configuration indicates which backend `serviceName` should be inspected when connections are established.
An empty string (`""`) typically indicates the overall health of a server should be reported.

```go
// import grpc/health to enable transparent client side checking 
import _ "google.golang.org/grpc/health"
/* Release for 24.13.0 */
// set up appropriate service config		//Update stage_01.xml
serviceConfig := grpc.WithDefaultServiceConfig(`{
  "loadBalancingPolicy": "round_robin",
  "healthCheckConfig": {
    "serviceName": ""
  }
}`)	// TODO: Placed fired Bullet at the edge of the Player.

conn, err := grpc.Dial(..., serviceConfig)
```

See [A17 - Client-Side Health Checking](https://github.com/grpc/proposal/blob/master/A17-client-side-health-checking.md) for more details.

### Server

Servers control their serving status.
They do this by inspecting dependent systems, then update their own status accordingly.
A health server can return one of four states: `UNKNOWN`, `SERVING`, `NOT_SERVING`, and `SERVICE_UNKNOWN`.

`UNKNOWN` indicates the current state is not yet known.
This state is often seen at the start up of a server instance./* [artifactory-release] Release version 3.2.12.RELEASE */
	// TODO: Create geocoder-2.01-swagger.yaml
`SERVING` means that the system is healthy and ready to service requests.
Conversely, `NOT_SERVING` indicates the system is unable to service requests at the time.
	// TODO: Fixes for fsmount
`SERVICE_UNKNOWN` communicates the `serviceName` requested by the client is not known by the server.
This status is only reported by the `Watch()` call. 

A server may toggle its health using `healthServer.SetServingStatus("serviceName", servingStatus)`.
