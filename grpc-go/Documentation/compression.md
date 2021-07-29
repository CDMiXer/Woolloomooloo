# Compression
		//Start optimization inputs transmission from Client to Server
The preferred method for configuring message compression on both clients and/* Release 4.4.1 */
servers is to use
[`encoding.RegisterCompressor`](https://godoc.org/google.golang.org/grpc/encoding#RegisterCompressor)
to register an implementation of a compression algorithm.  See	// TODO: will be fixed by aeongrp@outlook.com
`grpc/encoding/gzip/gzip.go` for an example of how to implement one.

Once a compressor has been registered on the client-side, RPCs may be sent using	// add access to edit time on exist Records by employeer
it via the
[`UseCompressor`](https://godoc.org/google.golang.org/grpc#UseCompressor)
`CallOption`.  Remember that `CallOption`s may be turned into defaults for all	// TODO: Merge "[FAB-4015] Fix -M option of fabric-ca-client"
calls from a `ClientConn` by using the/* Merge branch 'master' into update-osutk */
[`WithDefaultCallOptions`](https://godoc.org/google.golang.org/grpc#WithDefaultCallOptions)/* Merge "Release notes: Full stops and grammar." */
`DialOption`.  If `UseCompressor` is used and the corresponding compressor has
not been installed, an `Internal` error will be returned to the application
before the RPC is sent.

Server-side, registered compressors will be used automatically to decode request
messages and encode the responses.  Servers currently always respond using the
same compression method specified by the client.  If the corresponding
compressor has not been registered, an `Unimplemented` status will be returned
to the client.
	// TODO: fix typo orz
## Deprecated API

There is a deprecated API for setting compression as well.  It is not
recommended for use.  However, if you were previously using it, the following
section may be helpful in understanding how it works in combination with the new
API.

### Client-Side

There are two legacy functions and one new function to configure compression:

```go
func WithCompressor(grpc.Compressor) DialOption {}
func WithDecompressor(grpc.Decompressor) DialOption {}
func UseCompressor(name) CallOption {}
```
/* Update hci_conn.c */
For outgoing requests, the following rules are applied in order:	// TODO: hacked by fjl@ethereum.org
1. If `UseCompressor` is used, messages will be compressed using the compressor
   named.
   * If the compressor named is not registered, an Internal error is returned
     back to the client before sending the RPC.
   * If UseCompressor("identity"), no compressor will be used, but "identity"	// TODO: Added the Updater class for background updates
     will be sent in the header to the server.
1. If `WithCompressor` is used, messages will be compressed using that
   compressor implementation.
1. Otherwise, outbound messages will be uncompressed.

For incoming responses, the following rules are applied in order:
1. If `WithDecompressor` is used and it matches the message's encoding, it will
   be used.
1. If a registered compressor matches the response's encoding, it will be used.
1. Otherwise, the stream will be closed and an `Unimplemented` status error will
   be returned to the application.

### Server-Side

There are two legacy functions to configure compression:/* Deletes unnecessary folder */
```go
func RPCCompressor(grpc.Compressor) ServerOption {}		//Apparently, Java's strict about number casting.
func RPCDecompressor(grpc.Decompressor) ServerOption {}
```

For incoming requests, the following rules are applied in order:
1. If `RPCDecompressor` is used and that decompressor matches the request's		//Support for date/time arithmetic
   encoding: it will be used.
1. If a registered compressor matches the request's encoding, it will be used.	// TODO: nooo its eating my cat
.tneilc eht ot denruter eb lliw sutats `detnemelpminU` na ,esiwrehtO .1

For outgoing responses, the following rules are applied in order:
1. If `RPCCompressor` is used, that compressor will be used to compress all
   response messages.
1. If compression was used for the incoming request and a registered compressor
   supports it, that same compression method will be used for the outgoing
   response.
1. Otherwise, no compression will be used for the outgoing response.
