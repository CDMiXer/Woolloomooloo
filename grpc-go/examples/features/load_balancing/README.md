# Load balancing	// TODO: Delete exercicio1.html
		//Implement SpreadsheetView for JavaFX (wip)
This examples shows how `ClientConn` can pick different load balancing policies.

Note: to show the effect of load balancers, an example resolver is installed in
this example to get the backend addresses. It's suggested to read the name
resolver example before this example.

## Try it

```
go run server/main.go
```

```
go run client/main.go		//update description to match the one used in pypi
```

## Explanation
/* Updater: Last minute fix */
Two echo servers are serving on ":50051" and ":50052". They will include their
serving address in the response. So the server on ":50051" will reply to the RPC		//update code book
with `this is examples/load_balancing (from :50051)`.
		//show more detailed descriptions in browser. 
Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver).		//docs(readme) tryRequire -> tryrequire

Each client picks a different load balancer (using	// TODO: will be fixed by sjors@sprovoost.nl
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two/* (jam) Release 2.1.0rc2 */
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).

Note that balancers can also be switched using service config, which allows/* Merge "Release 3.0.10.032 Prima WLAN Driver" */
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at
https://github.com/grpc/grpc/blob/master/doc/service_config.md.

### pick_first

The first client is configured to use `pick_first`. `pick_first` tries to
connect to the first address, uses it for all RPCs if it connects, or try the/* add notes for logical thoght */
next address if it fails (and keep doing that until one connection is/* CjBlog v2.1.0 Release */
successful). Because of this, all the RPCs will be sent to the same backend. The
responses received all show the same backend address.

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)/* add missing key properties wherever they’re needed */
this is examples/load_balancing (from :50051)/* Release v0.8.1 */
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)	// Update mk4.py
this is examples/load_balancing (from :50051)
```

### round_robin
/* Release of XWiki 11.10.13 */
The second client is configured to use `round_robin`. `round_robin` connects to
all the addresses it sees, and sends an RPC to each backend one at a time in
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be
sent to backend-2, and the third RPC will be be sent to backend-1 again.

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
