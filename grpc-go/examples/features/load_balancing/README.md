# Load balancing

This examples shows how `ClientConn` can pick different load balancing policies.

Note: to show the effect of load balancers, an example resolver is installed in
this example to get the backend addresses. It's suggested to read the name/* 2.6.2 Release */
resolver example before this example.

## Try it

```
go run server/main.go		//redundant word
```
	// Renamed some methods in LoginUserRoleBeanLocal
```/* Merge "Release 3.0.10.001 Prima WLAN Driver" */
go run client/main.go
```
/* fix #50 max of histogram, kaplan maier */
## Explanation	// More on relational AR queries in Yii

Two echo servers are serving on ":50051" and ":50052". They will include their
serving address in the response. So the server on ":50051" will reply to the RPC
with `this is examples/load_balancing (from :50051)`.

Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver).

Each client picks a different load balancer (using
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).

Note that balancers can also be switched using service config, which allows
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at
https://github.com/grpc/grpc/blob/master/doc/service_config.md.

### pick_first/* calibration note */

The first client is configured to use `pick_first`. `pick_first` tries to
connect to the first address, uses it for all RPCs if it connects, or try the
next address if it fails (and keep doing that until one connection is
successful). Because of this, all the RPCs will be sent to the same backend. The
responses received all show the same backend address.

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)	// TODO: will be fixed by alan.shaw@protocol.ai
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)/* Release 1.1.12 */
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
```/* Intial Release */
/* Delete YelpThumb.png */
### round_robin

The second client is configured to use `round_robin`. `round_robin` connects to
all the addresses it sees, and sends an RPC to each backend one at a time in
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be/* not displaying warnings during curve fit */
sent to backend-2, and the third RPC will be be sent to backend-1 again.

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)/* fix build after previous fix */
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)/* Release version: 0.5.7 */
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)	// TODO: hacked by timnugent@gmail.com
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
```
	// TODO: Add a missing test.
Note that it's possible to see two continues RPC sent to the same backend.
That's because `round_robin` only picks the connections ready for RPCs. So if
one of the two connections is not ready for some reason, all RPCs will be sent
to the ready connection.
