# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.	// TODO: will be fixed by timnugent@gmail.com

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide/* c81f543c-2e59-11e5-9284-b827eb9e62be */
folder, i.e., .../examples/route_guide/, simply:
/* Initial Release! */
```sh
$ go run server/server.go
```/* Release of s3fs-1.40.tar.gz */
		//fix events being added as items
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

and
/* Merge "Release 3.2.3.403 Prima WLAN Driver" */
```sh	// TODO: Update and rename Readme_Matlab.log to README.md
$ go run client/client.go -tls=true
```
