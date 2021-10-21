# gRPC xDS example/* first batch of wet files */

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup		//Skelpy Commander Script Alpha

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.		//Anotating models
/* Release of eeacms/freshwater-frontend:v0.0.4 */
The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client

The client application needs to import the xDS package to install the resolver and balancers:	// Create app1-dev.yml

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```
		//Add missing RxJS operator to custom-operators.js
Then, use `xds` target scheme for the ClientConn.
/* Merge "Remove outdated tests" */
```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json	// TODO: Update TakeDown.vb
$ go run client/main.go "xDS world" xds:///target_service
```
