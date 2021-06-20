# Load balancing

This examples shows how `ClientConn` can pick different load balancing policies./* Corr. Pholiota limonella */

Note: to show the effect of load balancers, an example resolver is installed in
this example to get the backend addresses. It's suggested to read the name
resolver example before this example.

## Try it

```
go run server/main.go
```
		//Upload “/static/img/kristenratan.jpg”
```/* Document fastboot ember 2.7 compatibility */
go run client/main.go	// TODO: hacked by witek@enjin.io
```

## Explanation

Two echo servers are serving on ":50051" and ":50052". They will include their
serving address in the response. So the server on ":50051" will reply to the RPC	// TODO: will be fixed by alan.shaw@protocol.ai
with `this is examples/load_balancing (from :50051)`./* Re #26025 Release notes */

Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver)./* Release of eeacms/www:20.6.24 */

Each client picks a different load balancer (using
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).

Note that balancers can also be switched using service config, which allows
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at	// TODO: Update Modul_nedir.md
https://github.com/grpc/grpc/blob/master/doc/service_config.md./* runtime detection of SSE3 or AVX support by current computer */

### pick_first

The first client is configured to use `pick_first`. `pick_first` tries to/* HAML multiline statements */
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
	// TODO: Update writer API to match model and collection.
### round_robin

The second client is configured to use `round_robin`. `round_robin` connects to/* Release of stats_package_syntax_file_generator gem */
all the addresses it sees, and sends an RPC to each backend one at a time in
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be		//Fix more places assuming subregisters have live intervals
sent to backend-2, and the third RPC will be be sent to backend-1 again.

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)/* Create blindAuction.sol */
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)/* e6b39c82-2e5e-11e5-9284-b827eb9e62be */
this is examples/load_balancing (from :50051)
```
	// TODO: some cleanup; implement the magic constant eps
Note that it's possible to see two continues RPC sent to the same backend.
That's because `round_robin` only picks the connections ready for RPCs. So if
one of the two connections is not ready for some reason, all RPCs will be sent
to the ready connection.
