# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.
		//Fixed incorrect implicit name handling on empty root
Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information./* Merge "Release 1.0.0.173 QCACLD WLAN Driver" */

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide/* Add Release Notes for 1.0.0-m1 release */
folder, i.e., .../examples/route_guide/, simply:/* Create new class to represent DcosReleaseVersion (#350) */
		//Update README: Modified at parameter
```sh		//Bump version to 0.12.4
$ go run server/server.go
```
/* docs: 04 fix link */
Likewise, to run the client:/* Release of eeacms/forests-frontend:1.6.2 */

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true/* Block ui improvements. */
```
