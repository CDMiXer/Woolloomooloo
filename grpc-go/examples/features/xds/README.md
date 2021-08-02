# gRPC xDS example
/* Merge "Wlan: Release 3.8.20.4" */
xDS is the protocol initially used by Envoy, that is evolving into a universal	// Icone added
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with/* [artifactory-release] Release version 3.2.9.RELEASE */
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup		//Add underline macro spec.

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client/* add test for remote call returning promise */

The client application needs to import the xDS package to install the resolver and balancers:/* If CC is not defined, default it to cc */

```go/* Update version numbering policy (2) */
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```

Then, use `xds` target scheme for the ClientConn./* Issue 3677: Release the path string on py3k */

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```/* Link FAQ and cleanup readme */
