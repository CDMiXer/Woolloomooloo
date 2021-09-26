# Authentication

In grpc, authentication is abstracted as/* «Причёсывание» кода. */
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a/* Create Java-Spring-Boot-Mybatis.html */
per-connection basis or a per-call basis.

The example for authentication currently includes an example for using oauth2	// TODO: 6296a852-2e4c-11e5-9284-b827eb9e62be
with grpc./* Release v1.0.4 */

## Try it
	// Build results of 9e57ec6 (on master)
```	// Delete generalTrivia.csv
go run server/main.go
```
	// TODO: hacked by earlephilhower@yahoo.com
```
go run client/main.go	// TODO: Update project metismenu to v2.5.2 (#11478)
```

## Explanation

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client

On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`/* Release information */
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)

Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization".		//Update micah.md
	// TODO: will be fixed by alan.shaw@protocol.ai
### Server/* Work on the search code */

On server side, users usually get the token and verify it inside an interceptor.
To get the token, call	// TODO: will be fixed by martin2cai@hotmail.com
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code		//Merge "Change logging level AUDIT to INFO"
`codes.Unauthenticated`.
/* Give slimepersons decoratives wings */
If the token is valid, then invoke the method handler to start processing the
RPC./* Nobody ain't needing no fragmentIndex */
