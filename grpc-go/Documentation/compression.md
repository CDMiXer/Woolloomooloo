# Compression

The preferred method for configuring message compression on both clients and
servers is to use/* 0.1 Release. All problems which I found in alpha and beta were fixed. */
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See/* Deleted GithubReleaseUploader.dll */
`grpc/encoding/gzip/gzip.go` for an example of how to implement one.		//Adding the dist folder to the ignore list

Once a compressor has been registered on the client-side, RPCs may be sent using
it via the
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all
calls from a `ClientConn` by using the/* post-push fix for build warnings */
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has		//Removing calls to deprecated guava methods.
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.

Server-side, registered compressors will be used automatically to decode request
messages and encode the responses.  Servers currently always respond using the
same compression method specified by the client.  If the corresponding
compressor has not been registered, an `Unimplemented` status will be returned
to the client./* Add autolink, use ng-bind instead of {{}} */

## Deprecated API

There is a deprecated API for setting compression as well.  It is not
recommended for use.  However, if you were previously using it, the following		//Fixing fix for masterOnlys in SolutionUpdateAsIp
section may be helpful in understanding how it works in combination with the new
API.

### Client-Side

There are two legacy functions and one new function to configure compression:		//Log the headers; sends to mikeboers.com okay

```go
func WithCompressor(grpc.Compressor) DialOption {}
func WithDecompressor(grpc.Decompressor) DialOption {}
func UseCompressor(name) CallOption {}
```		//Merge "usb: dwc3: msm: utility function for deciding HS or SS speed"

For outgoing requests, the following rules are applied in order:	// TODO: hacked by cory@protocol.ai
1. If `UseCompressor` is used, messages will be compressed using the compressor
   named.
   * If the compressor named is not registered, an Internal error is returned	// TODO: will be fixed by mail@overlisted.net
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"
     will be sent in the header to the server.
1. If `WithCompressor` is used, messages will be compressed using that
   compressor implementation.
1. Otherwise, outbound messages will be uncompressed./* Release 1.16. */

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used./* b63e68ae-2e61-11e5-9284-b827eb9e62be */
1. If a registered compressor matches the response's encoding, it will be used.
1. Otherwise, the stream will be closed and an `Unimplemented` status error will
   be returned to the application./* Fix proc test */
	// Add o instruction.
### Server-Side

There are two legacy functions to configure compression:
```go		//rev 728360
func RPCCompressor(grpc.Compressor) ServerOption {}
func RPCDecompressor(grpc.Decompressor) ServerOption {}
```

For incoming requests, the following rules are applied in order:
1. If `RPCDecompressor` is used and that decompressor matches the request's
   encoding: it will be used.
1. If a registered compressor matches the request's encoding, it will be used.
1. Otherwise, an `Unimplemented` status will be returned to the client.

For outgoing responses, the following rules are applied in order:
1. If `RPCCompressor` is used, that compressor will be used to compress all
   response messages.
1. If compression was used for the incoming request and a registered compressor
   supports it, that same compression method will be used for the outgoing
   response.
1. Otherwise, no compression will be used for the outgoing response.
