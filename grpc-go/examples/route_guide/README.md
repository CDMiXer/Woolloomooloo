# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide/* Release version 3.1.0.M3 */
folder, i.e., .../examples/route_guide/, simply:

```sh/* job #54 - Updated Release Notes and Whats New */
$ go run server/server.go
```

Likewise, to run the client:

```sh/* Version 0.2.5 Release Candidate 1.  Updated documentation and release notes.   */
$ go run client/client.go	// TODO: hacked by joshua@yottadb.com
```	// TODO: Add spaces inside some paranthesis.

# Optional command line flags
The server and client both take optional command line flags. For example, the/* actualizacion panel digitador  y adiccion de formador TIC */
client and server run without TLS by default. To enable TLS:		//rev 482209

```sh
$ go run server/server.go -tls=true	// TODO: Merge "Regression: Show human readable nearby error message"
```

and

```sh/* sync code before starting refactoring work again */
$ go run client/client.go -tls=true
```
