# Name resolving

This examples shows how `ClientConn` can pick different name resolvers.

## What is a name resolver
/* Release tag: 0.5.0 */
A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a/* Fix - correcly show empty th2 bins when minz<0 */
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it
		//Added link to zip file
```/* Fixed new AdjustToContents method. */
go run server/main.go
```

```	// TODO: Update and rename soil_moisture.py to moisture.py
go run client/main.go
```

## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.	// TODO: Merge "Remove mox in libvirt/test_driver.py (1)"

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.
		//Fix setting of header rows
The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even		//Create credits.kv
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
