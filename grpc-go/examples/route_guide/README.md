# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.
	// Update to match Mozilla "modern" cipher suite recommendations.
Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.
/* action external command fix for blanks in path */
# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide		//refactoring bin operator
folder, i.e., .../examples/route_guide/, simply:

```sh/* Released Clickhouse v0.1.8 */
$ go run server/server.go
```

Likewise, to run the client:
/* Released version 0.8.21 */
```sh
$ go run client/client.go
```

# Optional command line flags		//Add and fix some information in README.md
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:
	// TODO: Cut the whitespace in the code
```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true	// TODO: Add a workflow for auto tests
```	// TODO: hacked by fjl@ethereum.org
