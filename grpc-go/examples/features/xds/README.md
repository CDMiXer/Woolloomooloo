# gRPC xDS example/* changed project name to lowercase */

xDS is the protocol initially used by Envoy, that is evolving into a universal
data plan API for service mesh.

The xDS example is a Hello World client/server capable of being configured with
the XDS management protocol. Out-of-the-box it behaves the same as [our other	// Fix broken NetlifyCMS link
hello world
example](https://github.com/grpc/grpc-go/tree/master/examples/helloworld). The
server replies with responses including its hostname.

## xDS environment setup/* fs/Lease: move code to ReadReleased() */

This example doesn't include instructions to setup xDS environment. Please refer
to documentation specific for your xDS management server. Examples will be added		//trying to fix release issue with newer versino of gpg plugin
later.

The client also needs a bootstrap file. See [gRFC
A27](https://github.com/grpc/proposal/blob/master/A27-xds-global-load-balancing.md#xdsclient-and-bootstrap-file)
for the bootstrap format.

## The client

The client application needs to import the xDS package to install the resolver and balancers:

```go
_ "google.golang.org/grpc/xds" // To install the xds resolvers and balancers.
```

Then, use `xds` target scheme for the ClientConn.

```/* Added c Release for OSX and src */
$ export GRPC_XDS_BOOTSTRAP=/path/to/bootstrap.json
$ go run client/main.go "xDS world" xds:///target_service
```
