# Authentication/* 45b537cc-2e4b-11e5-9284-b827eb9e62be */

As outlined in the [gRPC authentication guide](https://grpc.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.
	// TODO: will be fixed by nagydani@epointsystem.org
# Enabling TLS on a gRPC client

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```

# Enabling TLS on a gRPC server		//43478f4a-2e4a-11e5-9284-b827eb9e62be

```Go
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
if err != nil {	// TODO: will be fixed by witek@enjin.io
  log.Fatalf("Failed to generate credentials %v", err)
}
lis, err := net.Listen("tcp", ":0")
server := grpc.NewServer(grpc.Creds(creds))
...	// added video link to the readme
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication).

## Validating a token on the server
	// TODO: Show subtasks in delete confermation dialog again (was broken by r480)
Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)	// TODO: Added more asserts and tests.
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized/* Release 0.2.0 with corrected lowercase name. */
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New)./* Added support for deep merging overrides */

It is possible to configure token validation for all RPCs using an interceptor.
A server may configure either a
[grpc.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)	// TODO: Stricter deps.
or a
[grpc.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs	// Use wildcard dependency version.

To send an OAuth2 token with each RPC, a client may configure the
`grpc.DialOption`		//volt.1.4: Add missing constraint
[grpc.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpc.CallOption`
[grpc.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)
on each invocation of an RPC.

To create a `credentials.PerRPCCredentials`, use
[oauth.NewOauthAccess](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess).
Note, the OAuth2 implementation of `grpc.PerRPCCredentials` requires a client to use
[grpc.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)
to prevent any insecure transmission of tokens.

# Authenticating with Google	// TODO: Combine the channel tracker handler and the upstream handler
		//Rename Blade of the Land.txt to Blade of the Land
## Google Compute Engine (GCE)

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(oauth.NewComputeEngine()))
```

## JWT
/* Release of eeacms/www:18.9.4 */
```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)/* c5537ae0-2e49-11e5-9284-b827eb9e62be */
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))
```

