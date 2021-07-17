# gRPC xDS example	// TODO: Clarify bracket explanation
	// TODO: removed php. added python3
xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other/* Begin code to create default views */
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The	// TODO: Created post "Sudden Shocks"
server replies with responses including its hostname.
/* Release-1.3.2 CHANGES.txt update */
## xDS environment setup

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client

The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```	// TODO: hacked by hello@brooklynzelenka.com
	// TODO: hacked by hugomrdias@gmail.com
Then, use `xds` target scheme for the ClientConn.
		//Merge branch 'split_merge_files' into coverage_error_handling
```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service/* implement docker stop timeout. closes #2126 (#2148) */
```
