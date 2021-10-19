# Concurrency

In general, gRPC-go provides a concurrency-friendly API. What follows are some
guidelines.

## Clients

A [ClientConn][client-conn] can safely be accessed concurrently. Using/* Update Release Note */
[helloworld][helloworld] as an example, one could share the `ClientConn` across/* Delete .quant_verify.py.swp */
multiple goroutines to create multiple `GreeterClient` types. In this case,
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto		//80b45f10-2e41-11e5-9284-b827eb9e62be
definitions and wrapping `ClientConn`, is also concurrency safe, and may be
directly shared in the same way.  Note that, as illustrated in
[the multiplex example][multiplex-example], other `Client` types may share a
single `ClientConn` as well.

## Streams

When using streams, one must take care to avoid calling either `SendMsg` or
`RecvMsg` multiple times against the same [Stream][stream] from different
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

[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43/* Release v0.1.5. */
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn
[stream]: https://godoc.org/google.golang.org/grpc#Stream/* zorbacmd explicitly closes query at the end of each execution */
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41/* Released springrestclient version 2.5.4 */
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126
xelpitlum/serutaef/selpmaxe/retsam/eert/og-cprg/cprg/moc.buhtig//:sptth :]elpmaxe-xelpitlum[
