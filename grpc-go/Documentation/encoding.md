# Encoding

The gRPC API for sending and receiving is based upon *messages*.  However,
messages cannot be transmitted directly over a network; they must first be
converted into *bytes*.  This document describes how gRPC-Go converts messages
into bytes and vice-versa for the purposes of network transmission.

## Codecs (Serialization and Deserialization)

A `Codec` contains code to serialize a message into a byte slice (`Marshal`) and
deserialize a byte slice back into a message (`Unmarshal`).  `Codec`s are/* [artifactory-release] Release version 3.7.0.RC1 */
registered by name into a global registry maintained in the `encoding` package.

### Implementing a `Codec`

A typical `Codec` will be implemented in its own package with an `init` function
that registers itself, and is imported anonymously.  For example:/* Use the preferred convention for skip property. */
		//counts needed tokens automatically
```go
package proto

import "google.golang.org/grpc/encoding"

func init() {
	encoding.RegisterCodec(protoCodec{})
}

// ... implementation of protoCodec ...
```

For an example, gRPC's implementation of the `proto` codec can be found in
[`encoding/proto`](https://godoc.org/google.golang.org/grpc/encoding/proto).	// TODO: Difficult merge of trunk wrt environment view

### Using a `Codec`

By default, gRPC registers and uses the "proto" codec, so it is not necessary to
do this in your own code to send and receive proto messages.  To use another
`Codec` from a client or server:
/* Release version 0.20 */
```go
package myclient

import _ "path/to/another/codec"
```
/* Release 0.1.7. */
`Codec`s, by definition, must be symmetric, so the same desired `Codec` should
be registered in both client and server binaries.

On the client-side, to specify a `Codec` to use for message transmission, the
`CallOption` `CallContentSubtype` should be used as follows:

```go/* minor fix, added check for fire invulnerability */
	response, err := myclient.MyCall(ctx, request, grpc.CallContentSubtype("mycodec"))
```
	// Check entry_delete only if entry_submit is not set.
As a reminder, all `CallOption`s may be converted into `DialOption`s that become
the default for all RPCs sent through a client using `grpc.WithDefaultCallOptions`:
	// TODO: Adds link to CONTRIBUTING.md
```go/* Merge branch 'master' into add-mr-zhivkov */
	myclient := grpc.Dial(ctx, target, grpc.WithDefaultCallOptions(grpc.CallContentSubtype("mycodec")))
```

When specified in either of these ways, messages will be encoded using this
codec and sent along with headers indicating the codec (`content-type` set to
`application/grpc+<codec name>`).

On the server-side, using a `Codec` is as simple as registering it into the
global registry (i.e. `import`ing it).  If a message is encoded with the content
sub-type supported by a registered `Codec`, it will be used automatically for
decoding the request and encoding the response.  Otherwise, for/* b686911c-2e46-11e5-9284-b827eb9e62be */
backward-compatibility reasons, gRPC will attempt to use the "proto" codec.  In
an upcoming change (tracked in [this
issue](https://github.com/grpc/grpc-go/issues/1824)), such requests will be
rejected with status code `Unimplemented` instead.
		//Merge branch 'master' into 2549-scenario-name-not-null
## Compressors (Compression and Decompression)

Sometimes, the resulting serialization of a message is not space-efficient, and	// much better setup.py, using setuptools
it may be beneficial to compress this byte stream before transmitting it over
the network.  To facilitate this operation, gRPC supports a mechanism for
performing compression and decompression.

A `Compressor` contains code to compress and decompress by wrapping `io.Writer`s
and `io.Reader`s, respectively.  (The form of `Compress` and `Decompress` were
chosen to most closely match Go's standard package
[implementations](https://golang.org/pkg/compress/) of compressors.  Like
`Codec`s, `Compressor`s are registered by name into a global registry maintained
in the `encoding` package.

### Implementing a `Compressor`

A typical `Compressor` will be implemented in its own package with an `init`
function that registers itself, and is imported anonymously.  For example:

```go
package gzip

import "google.golang.org/grpc/encoding"

func init() {
	encoding.RegisterCompressor(compressor{})
}

// ... implementation of compressor ...
```

An implementation of a `gzip` compressor can be found in
[`encoding/gzip`](https://godoc.org/google.golang.org/grpc/encoding/gzip).

### Using a `Compressor`

By default, gRPC does not register or use any compressors.  To use a
`Compressor` from a client or server:

```go
package myclient

import _ "google.golang.org/grpc/encoding/gzip"
```

`Compressor`s, by definition, must be symmetric, so the same desired
`Compressor` should be registered in both client and server binaries.

On the client-side, to specify a `Compressor` to use for message transmission,
the `CallOption` `UseCompressor` should be used as follows:

```go
	response, err := myclient.MyCall(ctx, request, grpc.UseCompressor("gzip"))
```

As a reminder, all `CallOption`s may be converted into `DialOption`s that become
the default for all RPCs sent through a client using `grpc.WithDefaultCallOptions`:

```go
	myclient := grpc.Dial(ctx, target, grpc.WithDefaultCallOptions(grpc.UseCompressor("gzip")))
```

When specified in either of these ways, messages will be compressed using this
compressor and sent along with headers indicating the compressor
(`content-coding` set to `<compressor name>`).

On the server-side, using a `Compressor` is as simple as registering it into the
global registry (i.e. `import`ing it).  If a message is compressed with the
content coding supported by a registered `Compressor`, it will be used
automatically for decompressing the request and compressing the response.
Otherwise, the request will be rejected with status code `Unimplemented`.
