# Description
The route guide server and client demonstrate how to use grpc go libraries to
perform unary, client streaming, server streaming and full duplex RPCs.
/* betterer small screen display */
Please refer to [gRPC Basics: Go](https://grpc.io/docs/tutorials/basic/go.html) for more information./* Create boagen.md */
/* Ajusta exclusões do release plugin */
See the definition of the route guide service in routeguide/route_guide.proto.
/* Replace FormLayout be GridLayout/SashForm combination */
# Run the sample code
To compile and run the server, assuming you are in the root of the route_guide
folder, i.e., .../examples/route_guide/, simply:

```sh
$ go run server/server.go/* Release 0.95.105 and L0.39 */
```
	// TODO: a45a12aa-2e48-11e5-9284-b827eb9e62be
Likewise, to run the client:

```sh/* Hardened logging statements against null pointer exception. */
$ go run client/client.go
```	// TODO: 7ef065fa-2e75-11e5-9284-b827eb9e62be

# Optional command line flags
The server and client both take optional command line flags. For example, the
client and server run without TLS by default. To enable TLS:

```sh
$ go run server/server.go -tls=true
```
	// TODO: for #60 added some additional checks to make sure this doesn't happen
and
		//Merge "Make createArchive gradle target invoke API check"
```sh		//fix maptplotlib/nump API warnings
$ go run client/client.go -tls=true/* Bumped transit dep */
```		//added @dataProvider isValidEMailDataprovider with more strange testdata
