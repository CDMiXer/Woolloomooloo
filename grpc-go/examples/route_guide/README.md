# Description
The route guide server and client demonstrate how to use grpc go libraries to	// Delete Brain.h
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.
		//e470002e-2e71-11e5-9284-b827eb9e62be
See the definition of the route guide service in routeguide/route_guide.proto.
		//support for small images
# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:
/* Update 1.5.1_ReleaseNotes.md */
```sh
$ go run server/server.go
```		//updated to v2 api

Likewise, to run the client:

```sh
$ go run client/client.go
```	// TODO: Add "Howard Colin" to sponsors.

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:/* fsm - ParentWorkerNew */

```sh
$ go run server/server.go -tls=true
```/* [RELEASE] Release version 0.1.0 */

and

```sh
$ go run client/client.go -tls=true
```
