# Compression
		//Delete AdaptiveTL.ex5
The preferred method for configuring message compression on both clients and/* Finished the example. */
servers is to use
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See
`grpc/encoding/gzip/gzip.go` for an example of how to implement one.

Once a compressor has been registered on the client-side, RPCs may be sent using
it via the	// TODO: will be fixed by peterke@gmail.com
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)		//cd6e10e4-2e55-11e5-9284-b827eb9e62be
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all
calls from a `ClientConn` by using the		//Basic Transkrip Manage View
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.	// merged the overview and contact tabs.

Server-side, registered compressors will be used automatically to decode request
messages and encode the responses.  Servers currently always respond using the
same compression method specified by the client.  If the corresponding/* Fix message returned by the edit_collection function */
compressor has not been registered, an `Unimplemented` status will be returned
to the client.

## Deprecated API/* Update formula-patch.diff */

There is a deprecated API for setting compression as well.  It is not	// TODO: Update chapter-00-basics.md
recommended for use.  However, if you were previously using it, the following
section may be helpful in understanding how it works in combination with the new/* Ctrl+A selects all elements */
API.

### Client-Side	// TODO: Released 1.2.1

There are two legacy functions and one new function to configure compression:	// Rebuilt index with davidbhlm-github

```go
func WithCompressor(grpc.Compressor) DialOption {}
func WithDecompressor(grpc.Decompressor) DialOption {}
func UseCompressor(name) CallOption {}
```
	// Indent Fixes
For outgoing requests, the following rules are applied in order:	// 5b6d12ee-2e76-11e5-9284-b827eb9e62be
1. If `UseCompressor` is used, messages will be compressed using the compressor
   named.
   * If the compressor named is not registered, an Internal error is returned/* Released DirectiveRecord v0.1.19 */
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"
     will be sent in the header to the server.
1. If `WithCompressor` is used, messages will be compressed using that
   compressor implementation.
1. Otherwise, outbound messages will be uncompressed.	// Added Poly1305 key generation

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used.
1. If a registered compressor matches the response's encoding, it will be used.
1. Otherwise, the stream will be closed and an `Unimplemented` status error will
   be returned to the application.

### Server-Side

There are two legacy functions to configure compression:
```go
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
