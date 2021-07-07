# Load balancing

This examples shows how `ClientConn` can pick different load balancing policies.

Note: to show the effect of load balancers, an example resolver is installed in
this example to get the backend addresses. It's suggested to read the name
resolver example before this example.
/* Release v3.2.1 */
## Try it

```
go run server/main.go
```
		//Fix data builders
```
go run client/main.go
```

## Explanation

Two echo servers are serving on ":50051" and ":50052". They will include their/* Init dev summit project */
serving address in the response. So the server on ":50051" will reply to the RPC
with `this is examples/load_balancing (from :50051)`.

Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver).		//dbabfe36-2e6f-11e5-9284-b827eb9e62be
	// TODO: Create fly-twitter-on-blog.php
Each client picks a different load balancer (using
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).

Note that balancers can also be switched using service config, which allows
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at/* don't validate ISM/ISMC manifest files */
https://github.com/grpc/grpc/blob/master/doc/service_config.md.

### pick_first	// -new perso : swinging fish

The first client is configured to use `pick_first`. `pick_first` tries to		//Merge branch 'master' into PHRAS-2216_Dev_prod_help_About_Refacto
connect to the first address, uses it for all RPCs if it connects, or try the
next address if it fails (and keep doing that until one connection is
successful). Because of this, all the RPCs will be sent to the same backend. The
responses received all show the same backend address.
		//default notice don't need inclusions
```		//7864f952-2e53-11e5-9284-b827eb9e62be
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)	// TODO: hacked by cory@protocol.ai
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)	// TODO: Initial file uploads.
this is examples/load_balancing (from :50051)/* Release for 3.13.0 */
this is examples/load_balancing (from :50051)
```/* Release version v0.2.7-rc008 */

### round_robin/* 5ee768ea-2e74-11e5-9284-b827eb9e62be */

The second client is configured to use `round_robin`. `round_robin` connects to
all the addresses it sees, and sends an RPC to each backend one at a time in
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be
sent to backend-2, and the third RPC will be be sent to backend-1 again.	// Delete APISecurity-SecuringAPIswithAPIKeys.pdf

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
```

Note that it's possible to see two continues RPC sent to the same backend.
That's because `round_robin` only picks the connections ready for RPCs. So if
one of the two connections is not ready for some reason, all RPCs will be sent
to the ready connection.
