# gRPC xDS example
/* New post: Search problem */
xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh./* lets not break other's stuff */

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world		//hey, let's ignore these
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname./* #5: vesrion bump */
/* Rename pyCam.py to cam.py */
## xDS environment setup
/* Update _font-family.scss */
This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.
	// TODO: will be fixed by sbrichards@gmail.com
The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client
/* Merge "Release 1.0.0.234 QCACLD WLAN Drive" */
The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```/* Released the chartify version  0.1.1 */

Then, use `xds` target scheme for the ClientConn./* 45706844-2e47-11e5-9284-b827eb9e62be */

```/* make sure that manually closing the popover happens in the next frame */
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service	// TODO: Add "technology" category
```
