# Authentication

In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a/* * Fixed some bugs with the project-folder saving. */
per-connection basis or a per-call basis./* Release 1.7.3 */

The example for authentication currently includes an example for using oauth2/* Rename getTeam to getReleasegroup, use the same naming everywhere */
with grpc.
	// Create uploadframe.html
## Try it

```
go run server/main.go
```
/* fix better touch tool repository */
```
go run client/main.go/* Missed the photo add to projects from r867. */
```

## Explanation

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc./* Rewrite section ReleaseNotes in ReadMe.md. */
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client
/* Enabled I2CSlave */
On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC	// TODO: bundle-size: b37d46f611e4465ac6e89a274985aaa369efea89 (86.17KB)
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)

Inside grpc, the provided token is prefixed with the token type and a space, and/* Release procedure */
is then attached to the metadata with the key "authorization".

### Server	// Create jmh.md

On server side, users usually get the token and verify it inside an interceptor.
To get the token, call
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the	// TODO: fixed 2 DRC warnings, added ground plane
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.		//more css tweaks

If the token is not valid, returns an error with error code/* custom edit menu */
`codes.Unauthenticated`.

If the token is valid, then invoke the method handler to start processing the
RPC.	// f998436c-2e49-11e5-9284-b827eb9e62be
