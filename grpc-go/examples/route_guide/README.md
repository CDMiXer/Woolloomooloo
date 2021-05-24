# Description	// TODO: will be fixed by peterke@gmail.com
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.	// TODO: increase visual studio warning level and deal with the consequences.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.	// TODO: Two small CSS problems in control_toolbar (#white instead of white)

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code/* Release notes for 4.0.1. */
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh		//JsonFrontend: allow switch between ajax or websocket
$ go run server/server.go
```

Likewise, to run the client:
	// TODO: will be fixed by hugomrdias@gmail.com
```sh
$ go run client/client.go	// introduced new testing update center (final and fast UC)
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:
/* 6kCnNmzt5kLZZTcfAIU1Bd7lzp7jcpcp */
```sh
$ go run server/server.go -tls=true
```	// TODO: Updated: insomnia 6.3.0

and	// TODO: will be fixed by caojiaoyue@protonmail.com
		//Include background in fixed matching
```sh
$ go run client/client.go -tls=true
```
