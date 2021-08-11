# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.		//Use of Custom JLabel

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:/* Release 2.2.4 */

```sh
$ go run server/server.go
```/* Release redis-locks-0.1.2 */

Likewise, to run the client:

```sh
$ go run client/client.go
```
	// TODO: changed output folder to _subject_granular
# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:/* Removed fixed user color string */

```sh
$ go run server/server.go -tls=true
```
/* install only for Release build */
and
/* Make sure the stop casting button works */
```sh
$ go run client/client.go -tls=true
```
