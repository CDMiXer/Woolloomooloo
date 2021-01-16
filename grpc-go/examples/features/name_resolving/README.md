# Name resolving

This examples shows how `ClientConn` can pick different name resolvers.	// TODO: Merge "Speed up and reorganize rally jobs"

## What is a name resolver	// TODO: Native speaker sprinkles for quickstart.adoc

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a	// TODO: Create Shiny-Map Funktion (HTML)
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go/* 2LgmNf3dg6kCr3KonxXfyhBeHYTZzxQr */
```

```
go run client/main.go
```

## Explanation		//62418186-2f86-11e5-9241-34363bc765d8

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.
	// Cleaned up namespace (dependency, pd).
The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle/* implement Iterable.group and provide internal Map/Set implementation */
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will	// TODO: hacked by vyzo@hackzen.org
be created to the correct backend.
