# Name resolving

This examples shows how `ClientConn` can pick different name resolvers.		//removed duplicate columns for geocoding

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name/* MSO [Web Service - Tache #510] Couche DAO avec TU */
resolver is DNS.	// Merge "Added ephemeral disk limitationx"
	// TODO: [jgitflow-maven-plugin] updating poms for 1.4.16 branch with snapshot versions
In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it
		//Create InvUtils and transferFromToHand
```
go run server/main.go
```/* Add notification functions. */

```
go run client/main.go
```

## Explanation/* Add an option to amqp_publisher to use ssl for the connection. */

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See/* Updated Solution Files for Release 3.4.0 */
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`/* Release of eeacms/eprtr-frontend:0.0.2-beta.4 */
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.	// TODO: hacked by mail@overlisted.net
