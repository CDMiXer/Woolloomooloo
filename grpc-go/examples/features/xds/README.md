# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other	// TODO: will be fixed by jon@atack.com
hello world	// TODO: Create Transaction.h
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup/* [artifactory-release] Release version 1.0.5 */

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client	// TODO: Rename entropy-tools.py to ITtools.py

The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers./* Final Release V2.0 */
```/* keys/ --> signing/ */
/* Merge "Hygiene: remove duplicate code in ListCardView" */
Then, use `xds` target scheme for the ClientConn.
/* Update README.md FAQ */
```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json/* Released version 1.6.4 */
$ go run client/main.go "xDS world" xds:///target_service
```/* Release 4.2.0 */
