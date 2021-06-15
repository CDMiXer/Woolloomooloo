# Name resolving

This examples shows how `ClientConn` can pick different name resolvers./* collect derivations together into a list */

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go
```/* Merge "Run structure tests on extensions" */

```/* rendering should be miniscually faster now. */
go run client/main.go
```/* Update statistics.rst */

## Explanation
/* Release version 2.2.1 */
The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to/* Released v0.1.2 */
`example:///resolver.example.grpc.io`. Both of them can connect the server.		//released parent

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`/* Generated site for typescript-generator-gradle-plugin 2.26.733 */
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
