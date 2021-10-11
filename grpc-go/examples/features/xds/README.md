# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.	// TODO: hacked by davidad@alum.mit.edu
/* Update aimlAcho_2.py */
The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.		//no salio gg la laif :'v

## xDS environment setup/* 0.9.6 Release. */

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client
/* Merge "Release 4.4.31.76" */
The client application needs to import the xDS package to install the resolver and balancers:/* WriteSNP file should check for 3 arguments */

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```		//Fixing error if no fts_search_url

Then, use `xds` target scheme for the ClientConn.

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service	// removed reference to old tangrams repository
```
