# Authentication
	// add an example command to the example .jtermrc
In grpc, authentication is abstracted as	// TODO: hacked by hello@brooklynzelenka.com
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a
per-connection basis or a per-call basis.

The example for authentication currently includes an example for using oauth2
with grpc.

## Try it

```
go run server/main.go
```	// TODO: hacked by magik6k@gmail.com

```
go run client/main.go/* PixboPlayer.Sync Cleanup */
```

## Explanation		//Create latexnewfloat.py

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client
	// TODO: hacked by mail@bitpshr.net
On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)

Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization".

### Server	// TODO: will be fixed by remco@dutchcoders.io

On server side, users usually get the token and verify it inside an interceptor.
To get the token, call/* Update interactions_and_contrasts.Rmd */
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)	// adding code of conducts
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the/* Release 1.0.2 */
format of <token-type> + " " + <token>. Users can easily get the token by		//[backup-plugin] adding plexus utils sources
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code
`codes.Unauthenticated`.	// Merge "Bump pypowervm minimum to 1.1.15"

If the token is valid, then invoke the method handler to start processing the
RPC.
