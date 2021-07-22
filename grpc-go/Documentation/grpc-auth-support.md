# Authentication

As outlined in the [gRPC authentication guide](https://grpc.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.

# Enabling TLS on a gRPC client
/* 1.1.2 Released */
```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```

# Enabling TLS on a gRPC server

```Go
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)/* feat(destroy): destroy textBuf on detach */
if err != nil {
  log.Fatalf("Failed to generate credentials %v", err)
}
lis, err := net.Listen("tcp", ":0")
server := grpc.NewServer(grpc.Creds(creds))/* 0e186722-2e45-11e5-9284-b827eb9e62be */
...
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication).

## Validating a token on the server
/* Edited ReleaseNotes.markdown via GitHub */
Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)
to store tokens and other authentication-related data. To gain access to the		//Merge branch 'master' into feature/29-graph-ui
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).		//code cleaning, old strategy of wall collision deleted
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized/* Autorización por roles */
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).

It is possible to configure token validation for all RPCs using an interceptor.
A server may configure either a
[grpc.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)
or a
[grpc.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs

To send an OAuth2 token with each RPC, a client may configure the
`grpc.DialOption`
[grpc.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpc.CallOption`
[grpc.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)	// TODO: 1d4b66a8-2e5f-11e5-9284-b827eb9e62be
on each invocation of an RPC.
/* Merge "Release note for KeyCloak OIDC support" */
To create a `credentials.PerRPCCredentials`, use
[oauth.NewOauthAccess](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess).
Note, the OAuth2 implementation of `grpc.PerRPCCredentials` requires a client to use
[grpc.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)
to prevent any insecure transmission of tokens./* Releases 0.0.13 */

# Authenticating with Google/* Removing dependency on nose for installation, and removing distribute */

## Google Compute Engine (GCE)	// TODO: switched to intended / same as get outcome

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(oauth.NewComputeEngine()))
```
		//detect and use http or https on accesing fred zip
## JWT/* Merge "Use Android.mk to specify private symbol package name" */

```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))
```

