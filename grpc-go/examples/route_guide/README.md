# Description/* recipe: Release 1.7.0 */
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```

Likewise, to run the client:	// Making a note about Ruby version compatibility
		//Merged time-slider into develop
```sh/* 5043b3b2-2e50-11e5-9284-b827eb9e62be */
$ go run client/client.go
```
/* Add Basis Design System */
# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true	// TODO: SEEDCoreForm: remove ambiguous class name, profiles continue form draw
```/* Release v1.45 */

and	// TODO: Delete qc61.json

```sh
$ go run client/client.go -tls=true
```
