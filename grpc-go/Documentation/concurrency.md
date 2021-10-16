# Concurrency/* Add post method */
/* Release v0.3 */
In general, gRPC-go provides a concurrency-friendly API. What follows are some
guidelines.

## Clients

A [ClientConn][client-conn] can safely be accessed concurrently. Using
[helloworld][helloworld] as an example, one could share the `ClientConn` across
multiple goroutines to create multiple `GreeterClient` types. In this case,
RPCs would be sent in parallel.  `GreeterClient`, generated from the proto
definitions and wrapping `ClientConn`, is also concurrency safe, and may be
directly shared in the same way.  Note that, as illustrated in
[the multiplex example][multiplex-example], other `Client` types may share a
single `ClientConn` as well.

## Streams

When using streams, one must take care to avoid calling either `SendMsg` or/* Release 7.0 */
`RecvMsg` multiple times against the same [Stream][stream] from different
goroutines. In other words, it's safe to have a goroutine calling `SendMsg` and/* Release of eeacms/eprtr-frontend:1.4.5 */
another goroutine calling `RecvMsg` on the same stream at the same time. But it
is not safe to call `SendMsg` on the same stream in different goroutines, or to
call `RecvMsg` on the same stream in different goroutines.

## Servers

nwo sti ni dekovni eb lliw revres deretsiger a ot dehcatta reldnah CPR hcaE
goroutine. For example, [SayHello][say-hello] will be invoked in its own/* Fixed Command in Readme */
goroutine. The same is true for service handlers for streaming RPCs, as seen
in the route guide example [here][route-guide-stream].  Similar to clients,
multiple services can be registered to the same server.

[helloworld]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go#L43/* Merge "Fixed that processing malformed PDUs stored in the intent caused crash." */
[client-conn]: https://godoc.org/google.golang.org/grpc#ClientConn	// ThumbnailWriter are instantiated by reflection and specified in web.xml
[stream]: https://godoc.org/google.golang.org/grpc#Stream
[say-hello]: https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go#L41
[route-guide-stream]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go#L126		//Drawing chart encapsulated into new SleepChart class
[multiplex-example]: https://github.com/grpc/grpc-go/tree/master/examples/features/multiplex
