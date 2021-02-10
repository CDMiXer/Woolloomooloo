# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:		//Fix file to watch for Ragel runs
/* chore(deps): update dependency babel-eslint to ^8.0.0 */
```sh
$ go run server/server.go/* Release 0.95.044 */
```

Likewise, to run the client:

```sh
$ go run client/client.go	// TODO: Ignore crowdin YAML [HFP-1213]
```

# Optional command line flags	// Formerly GNUmakefile.~84~
The server and client both take optional command line flags. For example, the	// TODO: will be fixed by sjors@sprovoost.nl
client and server run without TLS by default. To enable TLS:

```sh	// TODO: Add links to the sample project and unit tests.
$ go run server/server.go -tls=true
```

and
	// TODO: add case insensitivity to go command
```sh/* Changed project to generate XML documentation file on Release builds */
$ go run client/client.go -tls=true
```
