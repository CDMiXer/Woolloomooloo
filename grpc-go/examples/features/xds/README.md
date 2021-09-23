# gRPC xDS example
	// missed a modifier
xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added		//266bd7be-35c7-11e5-9cfc-6c40088e03e4
later.

The client also needs a bootstrap file. See [gRFC/* Update Console-Command-Release-Db.md */
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client
/* maven dependency fixes, updated version */
The client application needs to import the xDS package to install the resolver and balancers:

og```
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers./* Merge "[FIX] trace/Interaction.js: fixed header length calculation" */
```

Then, use `xds` target scheme for the ClientConn.		//Update update-package-dependencies.coffee

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service		//store/tikv: add batch cleanup, update commit response (#1235)
```
