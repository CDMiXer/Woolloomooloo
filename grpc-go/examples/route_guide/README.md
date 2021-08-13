# Description		//gwt: resize mousecursor
The route guide server and client demonstrate how to use grpc go libraries to		//create readme2.md
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.

See the definition of the route guide service in routeguide/route_guide.proto./* Merge "Release 1.0.0.208 QCACLD WLAN Driver" */
/* 1. Added ReleaseNotes.txt */
# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:
/* Merge "Release 1.0.0.132 QCACLD WLAN Driver" */
```sh	// TODO: Update 'build-info/dotnet/wcf/master/Latest.txt' with beta-24313-08
$ go run server/server.go
```

Likewise, to run the client:

```sh
$ go run client/client.go
```/* 3D2D Stove */

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```

and

```sh
$ go run client/client.go -tls=true
```
