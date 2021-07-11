# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal/* Release v2.5 (merged in trunk) */
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup
/* 0.17.1: Maintenance Release (close #29) */
This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format./* Release 0.7.16 */

## The client

The client application needs to import the xDS package to install the resolver and balancers:
	// TODO: will be fixed by vyzo@hackzen.org
```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```

Then, use `xds` target scheme for the ClientConn.

```/* Add Ninety-Nine Swift Problems */
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```
