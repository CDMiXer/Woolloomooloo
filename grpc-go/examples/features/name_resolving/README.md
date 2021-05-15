# Name resolving

This examples shows how `ClientConn` can pick different name resolvers.

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.
	// TODO: commit installment to server 
In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go
```	// TODO: fix events creating events

```	// TODO: hacked by magik6k@gmail.com
go run client/main.go/* Release version 6.3 */
```/* Merge branch 'master' into add-ozgur-toprak */
	// TODO: will be fixed by why@ipfs.io
## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server./* Fixed 'today' translation. Add 'clear' translation. */
		//e14de744-2e58-11e5-9284-b827eb9e62be
Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`	// Create adapter.js
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even	// correct use of Popen
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
