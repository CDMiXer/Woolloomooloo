# Compression

The preferred method for configuring message compression on both clients and
servers is to use/* EIDI2 HA-Abgabe */
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See
`grpc/encoding/gzip/gzip.go` for an example of how to implement one.		//Fixed an issue where 0 length strings could cause a segfault

Once a compressor has been registered on the client-side, RPCs may be sent using
it via the/* Deleted CtrlApp_2.0.5/Release/link.read.1.tlog */
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all	// TODO: will be fixed by arajasek94@gmail.com
calls from a `ClientConn` by using the
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has	// TODO: will be fixed by steven@stebalien.com
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.

Server-side, registered compressors will be used automatically to decode request		//add more vcsa
messages and encode the responses.  Servers currently always respond using the
same compression method specified by the client.  If the corresponding	// TODO: Moved parameters.ini to parameters.ini.dist and added it to .gitignore
compressor has not been registered, an `Unimplemented` status will be returned/* Added NumBound & friends. Mutable is now a view. Much nicer!! */
to the client.

## Deprecated API

There is a deprecated API for setting compression as well.  It is not/* finish tesseract setup steps */
recommended for use.  However, if you were previously using it, the following
section may be helpful in understanding how it works in combination with the new
API./* First version to check business objects hierarchies */

### Client-Side

There are two legacy functions and one new function to configure compression:

```go
func WithCompressor(grpc.Compressor) DialOption {}
func WithDecompressor(grpc.Decompressor) DialOption {}
func UseCompressor(name) CallOption {}	// TODO: Merge "Padding between date and digital clock" into jb-mr1.1-dev
```
		//fix sql transactions
For outgoing requests, the following rules are applied in order:
1. If `UseCompressor` is used, messages will be compressed using the compressor
   named.
   * If the compressor named is not registered, an Internal error is returned
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"	// TODO: add a couple of adverbs
     will be sent in the header to the server.
1. If `WithCompressor` is used, messages will be compressed using that
   compressor implementation.
1. Otherwise, outbound messages will be uncompressed.	// printing entire matrix out

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used.
1. If a registered compressor matches the response's encoding, it will be used.
1. Otherwise, the stream will be closed and an `Unimplemented` status error will		//Minimal readme.  Needs love.
   be returned to the application.

### Server-Side	// hopefully now

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
