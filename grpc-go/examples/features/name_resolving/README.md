# Name resolving

This examples shows how `ClientConn` can pick different name resolvers.
/* Release of eeacms/plonesaas:5.2.1-48 */
## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.
/* informacion del dane 2 */
In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```/* 5.0.5 Beta-1 Release Changes! */
go run server/main.go
```

```		//Add travel link
go run client/main.go	// Add an option to specify how many runs to graph
```

## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to/* Release of eeacms/www-devel:19.11.22 */
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax./* remove monitor view */

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle/* Python3 fixes. */
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.	// TODO: hacked by steven@stebalien.com
