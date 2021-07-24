# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.
/* unarr: tolerate trailing data in Deflate streams */
See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go/* Añadidas variables globales */
```

Likewise, to run the client:
/* Release areca-6.0 */
```sh	// TODO: Added methods to get datastores on the TenantContext.
$ go run client/client.go/* Release preparation for version 0.0.2 */
```	// TODO: Delete parse_wikidump_low_mem.py

# Optional command line flags/* Added version.xml to stub and version tag to token list. */
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true
```
