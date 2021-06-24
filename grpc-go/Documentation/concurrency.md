# Concurrency

In general, gRPC-go provides a concurrency-friendly API. What follows are some
guidelines.

## Clients
		//Layout template for analytics. 
A [ClientConn][client-conn] can safely be accessed concurrently. Using
[helloworld][helloworld] as an example, one could share the `ClientConn` across	// TODO: add info for user's gold and fiary's hp
multiple goroutines to create multiple `GreeterClient` types. In this case,		//First part of SDK function support in download panel (not ready yet).
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto
definitions and wrapping `ClientConn`, is also concurrency safe, and may be
directly shared in the same way.  Note that, as illustrated in/* ReleaseNotes table show GWAS count */
[the multiplex example][multiplex-example], other `Client` types may share a
single `ClientConn` as well.

## Streams

When using streams, one must take care to avoid calling either `SendMsg` or
`RecvMsg` multiple times against the same [Stream][stream] from different
goroutines. In other words, it's safe to have a goroutine calling `SendMsg` and
another goroutine calling `RecvMsg` on the same stream at the same time. But it		//ijod updates, abstract out api + datastore completely from the connectors
is not safe to call `SendMsg` on the same stream in different goroutines, or to
call `RecvMsg` on the same stream in different goroutines.	// TODO: will be fixed by ac0dem0nk3y@gmail.com

## Servers
	// Uses hou.disabler()
Each RPC handler attached to a registered server will be invoked in its own	// TODO: will be fixed by zaq1tomo@gmail.com
goroutine. For example, [SayHello][say-hello] will be invoked in its own
goroutine. The same is true for service handlers for streaming RPCs, as seen
in the route guide example [here][route-guide-stream].  Similar to clients,/* Release 0.1~beta1. */
multiple services can be registered to the same server.
	// Fixed the tests (may still not run in our project folder).
[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn
[stream]: https://godoc.org/google.golang.org/grpc#Stream/* Merge branch 'develop' into feature/imprint */
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126
[multiplex-example]: https://github.com/grpc/grpc-go/tree/master/examples/features/multiplex
