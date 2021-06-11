# gRPC xDS example/* disable CI env var */

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The/* Edited middleware article */
server replies with responses including its hostname.
	// TODO: Text edit and cleanup
## xDS environment setup
	// TODO: hacked by joshua@yottadb.com
This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added
later.	// TODO: will be fixed by fjl@ethereum.org

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client

The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```/* Update wechat_little_app_controller.rb */
	// TODO: 7928738a-2e55-11e5-9284-b827eb9e62be
Then, use `xds` target scheme for the ClientConn.

```
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```
