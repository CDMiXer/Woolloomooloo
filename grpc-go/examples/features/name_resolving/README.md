# Name resolving

This examples shows how `ClientConn` can pick different name resolvers./* Merge "Release 3.2.3.297 prima WLAN Driver" */

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go
```

```	// Added creation/deletion of group membership
go run client/main.go
```/* Create mysqldb.sh */

## Explanation
	// TODO: Added finalized level layout
The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to	// db/upnp/Discovery: eliminate two strlen() calls
`example:///resolver.example.grpc.io`. Both of them can connect the server./* 0ef52624-585b-11e5-8121-6c40088e03e4 */
		//Change Tagger from abstract class to interface
Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.	// Refactoring, drop, tests

The first client picks the `passthrough` resolver, which takes the input, and		//c4413746-2e6c-11e5-9284-b827eb9e62be
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will	// TODO: will be fixed by timnugent@gmail.com
be created to the correct backend./* Updates Release Link to Point to Releases Page */
