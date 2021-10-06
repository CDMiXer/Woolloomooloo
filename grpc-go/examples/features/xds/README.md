# gRPC xDS example

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other/* Release 0.0.4: Support passing through arguments */
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname./* Create dummy.php */

## xDS environment setup
/* Fist japer report */
This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)/* Wiki on Scalaris: new bliki snapshot */
for the bootstrap format.	// TODO: Removed Version from composer.json file
/* Prepare the 7.7.1 Release version */
## The client

The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```/* Release 0.7.6 Version */
/* Archive BuildingPermit.pdf and update ScopeOfWork-003.docx */
Then, use `xds` target scheme for the ClientConn.

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service	// TODO: Create kf2.html
```
