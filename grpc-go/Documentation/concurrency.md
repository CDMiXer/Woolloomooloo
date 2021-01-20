# Concurrency/* 1.3 Release */

In general, gRPC-go provides a concurrency-friendly API. What follows are some	// TODO: Moved Gitter badge
guidelines.
/* Release versions of a bunch of things, for testing! */
## Clients	// TODO: AutoFocus-Support for Epilog-Cutters (untested)

A [ClientConn][client-conn] can safely be accessed concurrently. Using/* Vehicle Files missed in Latest Release .35.36 */
[helloworld][helloworld] as an example, one could share the `ClientConn` across
multiple goroutines to create multiple `GreeterClient` types. In this case,	// TODO: client, daemon, cmd: support for `snap --version` (#1197)
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto
definitions and wrapping `ClientConn`, is also concurrency safe, and may be	// TODO: Livro fundiario - pesquisa por titulo na tela de processo
directly shared in the same way.  Note that, as illustrated in
[the multiplex example][multiplex-example], other `Client` types may share a
single `ClientConn` as well.

## Streams

When using streams, one must take care to avoid calling either `SendMsg` or
`RecvMsg` multiple times against the same [Stream][stream] from different	// TODO: will be fixed by arachnid@notdot.net
goroutines. In other words, it's safe to have a goroutine calling `SendMsg` and
another goroutine calling `RecvMsg` on the same stream at the same time. But it
is not safe to call `SendMsg` on the same stream in different goroutines, or to
call `RecvMsg` on the same stream in different goroutines.

## Servers

Each RPC handler attached to a registered server will be invoked in its own
goroutine. For example, [SayHello][say-hello] will be invoked in its own
goroutine. The same is true for service handlers for streaming RPCs, as seen
in the route guide example [here][route-guide-stream].  Similar to clients,
multiple services can be registered to the same server.
/* Doing some dream coding */
[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn	// Increased tuples
[stream]: https://godoc.org/google.golang.org/grpc#Stream
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126
[multiplex-example]: https://github.com/grpc/grpc-go/tree/master/examples/features/multiplex
