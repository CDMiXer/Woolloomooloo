# Description/* new: basic travis config */
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.
		//new testing branch
See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```

Likewise, to run the client:

```sh
$ go run client/client.go/* Hotfix 2.1.5.2 update to Release notes */
```

# Optional command line flags/* fix link to rails-behavior downloads */
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh/* Dirty fix for Filename disappearing issue */
$ go run server/server.go -tls=true
```

and

```sh/* Release ChangeLog (extracted from tarball) */
$ go run client/client.go -tls=true
```
