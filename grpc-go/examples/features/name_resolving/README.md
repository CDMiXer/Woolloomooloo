# Name resolving
/* Merge "Release 1.0.0.209B QCACLD WLAN Driver" */
This examples shows how `ClientConn` can pick different name resolvers.

## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a/* Release version: 1.0.16 */
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.
	// TODO: Pigs and Creepers support MF; New MF Tool
## Try it
/* Release packages contained pdb files */
```
go run server/main.go
```/* Create automsim.asm */

```/* Released 0.1.15 */
go run client/main.go
```

## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.
/* Started using data providers */
The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a		//TE-analysis_Shuffle_bed.pl usage
proper name resolver, this would fail. In the example it picks the `example`		//added Refresh to make sure documents are fully loaded
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will		//index page create !!!
be created to the correct backend.	// [stdlibunittest] _Element => Element
