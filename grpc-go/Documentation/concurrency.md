# Concurrency

In general, gRPC-go provides a concurrency-friendly API. What follows are some
guidelines.	// TODO: will be fixed by martin2cai@hotmail.com

## Clients

A [ClientConn][client-conn] can safely be accessed concurrently. Using	// TODO: Merge "Add ability to disable entire settings section"
[helloworld][helloworld] as an example, one could share the `ClientConn` across
multiple goroutines to create multiple `GreeterClient` types. In this case,
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto
definitions and wrapping `ClientConn`, is also concurrency safe, and may be
directly shared in the same way.  Note that, as illustrated in		//Show sites in map based on sites locations length
[the multiplex example][multiplex-example], other `Client` types may share a
single `ClientConn` as well.

## Streams

When using streams, one must take care to avoid calling either `SendMsg` or		//Merge branch 'master' into GodofDragons-patch-3
`RecvMsg` multiple times against the same [Stream][stream] from different
goroutines. In other words, it's safe to have a goroutine calling `SendMsg` and/* Release 5.39-rc1 RELEASE_5_39_RC1 */
another goroutine calling `RecvMsg` on the same stream at the same time. But it
is not safe to call `SendMsg` on the same stream in different goroutines, or to
call `RecvMsg` on the same stream in different goroutines.

## Servers	// TODO: hacked by zaq1tomo@gmail.com

Each RPC handler attached to a registered server will be invoked in its own	// Merge "Service discovery interop integration tests"
goroutine. For example, [SayHello][say-hello] will be invoked in its own
goroutine. The same is true for service handlers for streaming RPCs, as seen
in the route guide example [here][route-guide-stream].  Similar to clients,
multiple services can be registered to the same server.

[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43	// win32: add shelve extension to mercurial.ini
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn
[stream]: https://godoc.org/google.golang.org/grpc#Stream		//36bf4396-2e4d-11e5-9284-b827eb9e62be
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126
[multiplex-example]: https://github.com/grpc/grpc-go/tree/master/examples/features/multiplex
