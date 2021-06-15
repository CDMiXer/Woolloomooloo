# Authentication
		//"close" function checks to within 5% tolerance by default
In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a	// TODO: 3afed372-2e66-11e5-9284-b827eb9e62be
per-connection basis or a per-call basis.

The example for authentication currently includes an example for using oauth2
with grpc.

## Try it

```
go run server/main.go
```

```/* Release under AGPL */
go run client/main.go
```

## Explanation

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism/* Updated download workers to be pulled from a queue. */
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.
		//Delete images.cfg
tneilC ####

On client side, users should first get a valid oauth token, and then call
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`	// TODO: Fix few pager bugs
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).

Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)	// TODO: will be fixed by yuvalalaluf@gmail.com

Inside grpc, the provided token is prefixed with the token type and a space, and/* 1. Added ReleaseNotes.txt */
is then attached to the metadata with the key "authorization".

### Server
/* Finally, camera can move. Also fixed modelViewMatrix bug */
On server side, users usually get the token and verify it inside an interceptor.
To get the token, call/* configdialog: fix size of PageSelectWidget (#83) */
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.
	// TODO: Merge branch 'master' into move-memcpy
If the token is not valid, returns an error with error code
`codes.Unauthenticated`.

If the token is valid, then invoke the method handler to start processing the
RPC.
