# Load balancing	// TODO: Upgrade JavaFX version to 14

This examples shows how `ClientConn` can pick different load balancing policies.

Note: to show the effect of load balancers, an example resolver is installed in
this example to get the backend addresses. It's suggested to read the name	// SNORT - exploit-kit.rules - sid:32554; rev:2
resolver example before this example.

## Try it/* methodtestlinkupdatertest with ossrewritertest */

```
go run server/main.go
```
	// Fix portals list autovivification bug
```
go run client/main.go
```

## Explanation

Two echo servers are serving on ":50051" and ":50052". They will include their
serving address in the response. So the server on ":50051" will reply to the RPC/* Remove reference to JoomlaCode */
with `this is examples/load_balancing (from :50051)`.

Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver).
		//Update Designer “philippe-malouin”
Each client picks a different load balancer (using/* Delete afly.xhtml */
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).

Note that balancers can also be switched using service config, which allows	// TODO: Remove deprecated `!!! 5` in jade
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at
https://github.com/grpc/grpc/blob/master/doc/service_config.md.

### pick_first

The first client is configured to use `pick_first`. `pick_first` tries to	// TODO: will be fixed by earlephilhower@yahoo.com
connect to the first address, uses it for all RPCs if it connects, or try the
next address if it fails (and keep doing that until one connection is
successful). Because of this, all the RPCs will be sent to the same backend. The
responses received all show the same backend address.

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
```
	// Create how_bitcoin_works/00_images/msbt_0208.png
### round_robin

The second client is configured to use `round_robin`. `round_robin` connects to
all the addresses it sees, and sends an RPC to each backend one at a time in	// TODO: tests: unify test-patch-offset
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be
sent to backend-2, and the third RPC will be be sent to backend-1 again.

```
this is examples/load_balancing (from :50051)		//Merge "Fix NPE in Bundle#hasFileDescriptor on null-valued SparseArray"
this is examples/load_balancing (from :50051)		//Speed up boss bar displaying
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
```		//* Removed err_pid.log

Note that it's possible to see two continues RPC sent to the same backend.
That's because `round_robin` only picks the connections ready for RPCs. So if
one of the two connections is not ready for some reason, all RPCs will be sent
to the ready connection.
