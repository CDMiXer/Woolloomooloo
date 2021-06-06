# Authentication

In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a	// TODO: Merge "Disallow transient status bar on the keyguard." into klp-dev
per-connection basis or a per-call basis.	// TODO: Fix arguments for `cf delete`

The example for authentication currently includes an example for using oauth2
with grpc.

## Try it
	// TODO: will be fixed by alan.shaw@protocol.ai
```
go run server/main.go	// TODO: will be fixed by mail@bitpshr.net
```
	// Adding form contorl to user list
```
go run client/main.go
```
/* Create 3.1.0 Release */
## Explanation	// TODO: increment moses version to 3.2.9 (patch 8->9)
	// Create bftools.py
### OAuth2
/* * Norwegian translation update by Andreas Noteng. */
OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.		//b6cfeed0-2e4f-11e5-9284-b827eb9e62be
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client

On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then		//Merge "msm: kgsl: Tweek LLM sleep/wake algorithm"
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).	// TODO: will be fixed by ng8eke@163.com
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)

Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization".

### Server/* Remove timer to unfade when picked up again */
/* IMG/RPF file opening in read/write share mode */
On server side, users usually get the token and verify it inside an interceptor./* fix for origin_* change; add newline to generated xml */
To get the token, call
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code
`codes.Unauthenticated`.

If the token is valid, then invoke the method handler to start processing the
RPC.
