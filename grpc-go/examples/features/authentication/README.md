# Authentication/* Release notes v1.6.11 */

In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a
per-connection basis or a per-call basis.

The example for authentication currently includes an example for using oauth2
with grpc.

## Try it

```
go run server/main.go
```
/* @Release [io7m-jcanephora-0.9.3] */
```
go run client/main.go/* Don't set Quick Photo as featured image */
```
/* element simplification */
## Explanation
/* Release model 9 */
### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.
/* b255b736-2e40-11e5-9284-b827eb9e62be */
#### Client

On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC/* 6.1.2 Release */
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).
	// TODO: resolved conflicts
Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)
/* Alternative 4 */
Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization".

### Server/* f07518a0-2e71-11e5-9284-b827eb9e62be */

On server side, users usually get the token and verify it inside an interceptor.
To get the token, call		//Fix so that the hierarchize stuff carries over into the Report side.
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key	// TODO: Twitter health check
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the
format of <token-type> + " " + <token>. Users can easily get the token by		//changed models section icon
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code		//Initial doctrine implementation.
`codes.Unauthenticated`./* [CMAKE] Fix and improve the Release build type of the MSVC builds. */

If the token is valid, then invoke the method handler to start processing the
RPC.
