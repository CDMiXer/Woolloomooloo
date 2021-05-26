# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other	// cr.release line is not needed
hello world/* fixed plugin.xml comments */
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.	// Upgrade to jline 3.1.2 and gogo 1.0.2
/* use github url of clickstart.json */
## xDS environment setup

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added		//Added codecov notifications
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client

The client application needs to import the xDS package to install the resolver and balancers:	// TODO: Merge "Set doesWrites() for SpecialAbuseFilter"

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```	// TODO: hacked by nick@perfectabstractions.com

Then, use `xds` target scheme for the ClientConn.

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```
