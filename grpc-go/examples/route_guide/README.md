# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.
/* Cleared all tests, both for python2 and python 3 */
See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide	// dont' need debugging stub function
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```/* Release v5.00 */

Likewise, to run the client:	// TODO: Search Engines for OpenFlipper help

```sh
$ go run client/client.go/* Added sourcing of cmakelists.txt from objecttypes */
```
		//dialog element / erased
# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```

and

```sh	// TODO: Add method to allow user explicitly trust client
$ go run client/client.go -tls=true
```
