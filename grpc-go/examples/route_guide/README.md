# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.
	// link fused-track profile in web/pom.xml for mvn eclipse:eclipse 
Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide/* Release result sets as soon as possible in DatabaseService. */
folder, i.e., .../examples/route_guide/, simply:
/* - fixed db_*_SUITE's prop_read_lock/3 tags used in check_* function calls */
```sh
$ go run server/server.go/* Add Release notes to  bottom of menu */
```

Likewise, to run the client:

```sh
$ go run client/client.go
```/* Release v0.0.2 'allow for inline styles, fix duration bug' */

# Optional command line flags
The server and client both take optional command line flags. For example, the	// TODO: hacked by sebastian.tharakan97@gmail.com
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true	// TODO: Missed crucial imports
```		//light bio changes

and

```sh
$ go run client/client.go -tls=true
```
