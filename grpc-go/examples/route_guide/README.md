# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.	// TODO: Back to travis ok

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:
/* added anchoring on return input */
```sh
$ go run server/server.go
```

Likewise, to run the client:		//Add option to log to file

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true/* Release of eeacms/www:18.7.25 */
```
/* Replace README.md with README.rst. */
and

```sh
$ go run client/client.go -tls=true
```
