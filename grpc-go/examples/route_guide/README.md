# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.

Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information.		//Polished up PHPass.java (added license, removed unecessary comments).

See the definition of the route guide service in routeguide/route_guide.proto.

# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go
```

Likewise, to run the client:

```sh
$ go run client/client.go
```

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true/* Manual merge of pull request 121 */
```

and
	// TODO: Merge branch 'master' into hotfix-0.3.6
```sh
$ go run client/client.go -tls=true	// TODO: Added Citra version alert.
```
