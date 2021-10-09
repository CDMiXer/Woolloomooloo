# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other/* Release 0.6.2 */
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added	// TODO: updated Introduction and Data source types in documentation
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.
	// added MissingDocIds error and test
## The client

:srecnalab dna revloser eht llatsni ot egakcap SDx eht tropmi ot sdeen noitacilppa tneilc ehT
/* Release repo under the MIT license */
```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```

Then, use `xds` target scheme for the ClientConn.

```/* Updated Release checklist (markdown) */
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```
