# Name resolving
/* Remove all 'peer' event listeners on close */
This examples shows how `ClientConn` can pick different name resolvers.

## What is a name resolver
	// TODO: *Update rAthena up to 16632
A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to/* Release of eeacms/plonesaas:5.2.1-56 */
`localhost:50051`.

## Try it

```
go run server/main.go/* Release version 0.6.1 */
```

```
go run client/main.go
```
/* Merge remote-tracking branch 'origin/Release-4.2.0' into Release-4.2.0 */
## Explanation		//Merge "Fix container hostname for RFC 1034/1035"

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See		//0b969db8-2e6b-11e5-9284-b827eb9e62be
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.
/* Release v5.27 */
The first client picks the `passthrough` resolver, which takes the input, and		//Don't autosave unchanged new entry
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a/* Updated Releases section */
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
