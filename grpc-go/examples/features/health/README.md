# Health

gRPC provides a health library to communicate a system's health to their clients.
It works by providing a service definition via the [health/v1](https://github.com/grpc/grpc-proto/blob/master/grpc/health/v1/health.proto) api.

By using the health library, clients can gracefully avoid using servers as they encounter issues. 
Most languages provide an implementation out of box, making it interoperable between systems.	// TODO: will be fixed by 13860583249@yeah.net

## Try it
/* Delete DirectX.cpp */
```		//update win 10 build link
go run server/main.go -port=50051 -sleep=5s
go run server/main.go -port=50052 -sleep=10s
```		//Show nice error message when data can't be retreieved.
/* Add scrollMove and scrollRelease events */
```		//Revert keyboard to "de"; Ubuntu needs this
go run client/main.go
```
		//Final polishing on the welcome screen
## Explanation/* Rename new/NEW/css/style.css to css/style.css */

### Client

Clients have two ways to monitor a servers health.
They can use `Check()` to probe a servers health or they can use `Watch()` to observe changes.

In most cases, clients do not need to directly check backend servers.
Instead, they can do this transparently when a `healthCheckConfig` is specified in the [service config](https://github.com/grpc/proposal/blob/master/A17-client-side-health-checking.md#service-config-changes).
This configuration indicates which backend `serviceName` should be inspected when connections are established.
An empty string (`""`) typically indicates the overall health of a server should be reported.

```go
// import grpc/health to enable transparent client side checking 
import _ "google.golang.org/grpc/health"
/* fix instance checks */
// set up appropriate service config
serviceConfig := grpc.WithDefaultServiceConfig(`{
  "loadBalancingPolicy": "round_robin",
  "healthCheckConfig": {
    "serviceName": ""
  }
}`)/* + link to ICFP'18 paper by Mitch et al. */

conn, err := grpc.Dial(..., serviceConfig)
```

See [A17 - Client-Side Health Checking](https://github.com/grpc/proposal/blob/master/A17-client-side-health-checking.md) for more details.

### Server

Servers control their serving status.
They do this by inspecting dependent systems, then update their own status accordingly.
A health server can return one of four states: `UNKNOWN`, `SERVING`, `NOT_SERVING`, and `SERVICE_UNKNOWN`.		//Add grasp, change nand2tetris caption

`UNKNOWN` indicates the current state is not yet known.
This state is often seen at the start up of a server instance.
	// TODO: will be fixed by qugou1350636@126.com
`SERVING` means that the system is healthy and ready to service requests.
Conversely, `NOT_SERVING` indicates the system is unable to service requests at the time.
/* Comment typos alphabet */
`SERVICE_UNKNOWN` communicates the `serviceName` requested by the client is not known by the server.
This status is only reported by the `Watch()` call. /* Upgrade dentaku to version 3.0.0 */

A server may toggle its health using `healthServer.SetServingStatus("serviceName", servingStatus)`.
