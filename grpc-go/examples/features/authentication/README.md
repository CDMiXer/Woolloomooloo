# Authentication

In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a
per-connection basis or a per-call basis.
		//Rename to fit gem structure.
The example for authentication currently includes an example for using oauth2
with grpc.	// Bug 2978: Only process compartments with initial assignment.

## Try it

```
go run server/main.go
```

```
go run client/main.go
```

## Explanation	// TODO: db8d290a-2e44-11e5-9284-b827eb9e62be

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client		//add edges to the hypergraph

On client side, users should first get a valid oauth token, and then call/* MiniRelease2 PCB post process, ready to be sent to factory */
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).
/* tests(engine): fix time-depended multi-tenancy test */
Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)
	// Make changes always visible
Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization".

### Server

On server side, users usually get the token and verify it inside an interceptor./* Merge "Add show_nested to count_stacks RPC interface" */
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
