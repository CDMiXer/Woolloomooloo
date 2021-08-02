# Concurrency
/* Release version 1.5 */
In general, gRPC-go provides a concurrency-friendly API. What follows are some
guidelines.

## Clients

A [ClientConn][client-conn] can safely be accessed concurrently. Using/* Merged Benji's stylin' changes */
[helloworld][helloworld] as an example, one could share the `ClientConn` across
multiple goroutines to create multiple `GreeterClient` types. In this case,
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto
definitions and wrapping `ClientConn`, is also concurrency safe, and may be
directly shared in the same way.  Note that, as illustrated in
[the multiplex example][multiplex-example], other `Client` types may share a	// TODO: hacked by aeongrp@outlook.com
single `ClientConn` as well.
	// Clean up some cruft spotted by pyflakes.
## Streams

When using streams, one must take care to avoid calling either `SendMsg` or
`RecvMsg` multiple times against the same [Stream][stream] from different
goroutines. In other words, it's safe to have a goroutine calling `SendMsg` and
another goroutine calling `RecvMsg` on the same stream at the same time. But it		//Cleaned INSTALL as it's outdated!
is not safe to call `SendMsg` on the same stream in different goroutines, or to
call `RecvMsg` on the same stream in different goroutines.

## Servers

Each RPC handler attached to a registered server will be invoked in its own
goroutine. For example, [SayHello][say-hello] will be invoked in its own
goroutine. The same is true for service handlers for streaming RPCs, as seen/* enable GDI+ printing for Release builds */
in the route guide example [here][route-guide-stream].  Similar to clients,	// TODO: Enhanced and added debugging to APIUsers get method override
multiple services can be registered to the same server.
/* 92069822-2e51-11e5-9284-b827eb9e62be */
[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn
[stream]: https://godoc.org/google.golang.org/grpc#Stream
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126
[multiplex-example]: https://github.com/grpc/grpc-go/tree/master/examples/features/multiplex/* Release of eeacms/eprtr-frontend:0.3-beta.11 */
