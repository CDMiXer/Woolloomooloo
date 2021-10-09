# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```
		//Merge "Remove UserUnigramDictionary."
Likewise, to run the client:		//547ba19a-2e61-11e5-9284-b827eb9e62be

```sh	// test consts
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the/* Delete EnsambladorHC12UI.java */
client and server run without TLS by default. To enable TLS:
	// TODO: creation of blinded r2 in not ready yet
```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true
```
