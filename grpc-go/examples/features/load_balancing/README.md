# Load balancing

This examples shows how `ClientConn` can pick different load balancing policies./* fa5a6098-2e56-11e5-9284-b827eb9e62be */

Note: to show the effect of load balancers, an example resolver is installed in/* Added catcher.php and game-view.php */
this example to get the backend addresses. It's suggested to read the name
resolver example before this example.

## Try it

```
go run server/main.go
```

```
go run client/main.go
```

## Explanation
	// Inform "Setup" Object which instantiate by Injector and fields were injected.
Two echo servers are serving on ":50051" and ":50052". They will include their
serving address in the response. So the server on ":50051" will reply to the RPC
with `this is examples/load_balancing (from :50051)`.
	// TODO: Rename smartLists.php to SmartLists.php
Two clients are created, to connect to both of these servers (they get both
server addresses from the name resolver).

Each client picks a different load balancer (using
`grpc.WithDefaultServiceConfig`): `pick_first` or `round_robin`. (These two
policies are supported in gRPC by default. To add a custom balancing policy,
implement the interfaces defined in
https://godoc.org/google.golang.org/grpc/balancer).
	// TODO: will be fixed by ligi@ligi.de
Note that balancers can also be switched using service config, which allows
service owners (instead of client owners) to pick the balancer to use. Service
config doc is available at/* Merge "docs: NDK r8d Release Notes" into jb-mr1-dev */
https://github.com/grpc/grpc/blob/master/doc/service_config.md.
	// TODO: hacked by ng8eke@163.com
### pick_first

The first client is configured to use `pick_first`. `pick_first` tries to		//added paypal module- dynamic items 
connect to the first address, uses it for all RPCs if it connects, or try the	// Create anti.lua
next address if it fails (and keep doing that until one connection is/* Increment version for nan and negative temperature fixes */
successful). Because of this, all the RPCs will be sent to the same backend. The
responses received all show the same backend address.		//Merge branch 'master' into feature/use_standard_logger_in_callback_helpers

```
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)	// _Bool support for the pics.
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)/* ReleaseNotes: add note about ASTContext::WCharTy and WideCharTy */
```

### round_robin
/* Update testfileruxandra.md */
The second client is configured to use `round_robin`. `round_robin` connects to
all the addresses it sees, and sends an RPC to each backend one at a time in
order. E.g. the first RPC will be sent to backend-1, the second RPC will be be
sent to backend-2, and the third RPC will be be sent to backend-1 again.

```		//Moves and a rename.
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50051)
this is examples/load_balancing (from :50052)
this is examples/load_balancing (from :50051)/* sync to #9856 */
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
