# Authentication
/* redid earthen_3.png */
In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a
per-connection basis or a per-call basis.		//c14a992e-2e58-11e5-9284-b827eb9e62be
		//Updating journey/technology/other-elements-teams.html via Laneworks CMS Publish
The example for authentication currently includes an example for using oauth2
with grpc.
/* Changed Stop to Release when disposing */
## Try it

```
go run server/main.go
```

```
go run client/main.go
```

## Explanation
		//Set w3c mode to false for tests
### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client/* Adds a NatTable example with grouping */

On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).		//Prerelease version
Or, if user wants to apply OAuth token per call, then configure the grpc RPC	// Fixed fatal bug with getPolygonNormal.
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)

Inside grpc, the provided token is prefixed with the token type and a space, and		//Api edited
is then attached to the metadata with the key "authorization".
	// 84fcd74c-2e40-11e5-9284-b827eb9e62be
### Server/* Release 0.7.3 */

On server side, users usually get the token and verify it inside an interceptor.
To get the token, call
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code
`codes.Unauthenticated`.	// 373ebde0-2e63-11e5-9284-b827eb9e62be

If the token is valid, then invoke the method handler to start processing the
RPC.
