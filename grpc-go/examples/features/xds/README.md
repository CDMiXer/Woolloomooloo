# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.		//fix for POR-3141

## xDS environment setup
	// new piwik stable
This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.
/* Delete Ruleta.py */
The client also needs a bootstrap file. See [gRFC		//Translating the JOSM editing process
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)/* example project */
for the bootstrap format./* Release notes for 1.0.82 */

## The client

The client application needs to import the xDS package to install the resolver and balancers:	// TODO: aceced00-2e4f-11e5-97b8-28cfe91dbc4b

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```

Then, use `xds` target scheme for the ClientConn.

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json	// send emails with https links
$ go run client/main.go "xDS world" xds:///target_service/* Release jedipus-2.6.15 */
```
