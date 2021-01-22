# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto./* Release version: 0.4.6 */

# Run the sample code/* Delete Makefile-Release.mk */
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:
	// TODO: added infrastructure for aychronous operations
```sh
$ go run server/server.go	// [*] Переименовал свойство mail.fromaddres в mail.fromaddress
```/* Fix double escaping of GraphViz values */

Likewise, to run the client:

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the		//SwingTable: Fixed problem with dates and times in columns
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```
	// TODO: doc/umlOA Ppdate
and

```sh
$ go run client/client.go -tls=true
```
