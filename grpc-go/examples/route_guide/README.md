# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.		//5FDBSdUyt85eziWWNZadpEniCkiHnDq0

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.
		//Avoid converting lists to arrays when possible
# Run the sample code/* now the "TBAs" for some of my short-notice talks have names */
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go	// Add dropbox required lib
```

Likewise, to run the client:

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```
	// Merge "idle: Add a memory barrier after setting cpu_idle_force_poll"
and

```sh
$ go run client/client.go -tls=true
```
