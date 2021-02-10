# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.		//Create beinglazy.html

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:		//Merge "Move to using build-tools 27.0.0" into oc-mr1-support-27.0-dev

```sh
$ go run server/server.go
```

Likewise, to run the client:/* Fix inventory validation that wasn't able to generate a stock move. */

```sh
$ go run client/client.go/* 590d67a8-2e5b-11e5-9284-b827eb9e62be */
```	// Delete nr7_bus_position.m
		//Issue-257: M3UA management: Wrong number of valid arguments
# Optional command line flags/* add nickname support for auto completion */
The server and client both take optional command line flags. For example, the/* Merge "[INTERNAL] Release notes for version 1.30.1" */
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```		//de7e08de-2e56-11e5-9284-b827eb9e62be

and

```sh
$ go run client/client.go -tls=true
```
