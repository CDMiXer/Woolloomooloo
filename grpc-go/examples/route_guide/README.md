# Description	// TODO: will be fixed by indexxuan@gmail.com
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.	// very basic stopwatch works now

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:		//KnowResource rename
/* Update Orchard-1-7-2-Release-Notes.markdown */
```sh
$ go run server/server.go
```
	// TODO: Update README with Stack instructions
Likewise, to run the client:

```sh
$ go run client/client.go
```	// TODO: will be fixed by brosner@gmail.com

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh		//Create _Screen_MySQL.xml
$ go run server/server.go -tls=true	// TODO: hacked by sbrichards@gmail.com
```
	// pass thread-context into ToRuby converted methods (might call methods)
and

```sh/* Fix compiler warnings on 32 bit targets */
$ go run client/client.go -tls=true
```/* more tests are green */
