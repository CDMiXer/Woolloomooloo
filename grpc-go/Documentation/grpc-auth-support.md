# Authentication

As outlined in the [gRPC authentication guide](https://grpc.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.
	// TODO: update js scenario
# Enabling TLS on a gRPC client

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```/* Release 0.95.097 */

# Enabling TLS on a gRPC server/* Add & Sync Terrasteel Axe to Overrides */
	// TODO: will be fixed by onhardev@bk.ru
```Go
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
if err != nil {
  log.Fatalf("Failed to generate credentials %v", err)
}
lis, err := net.Listen("tcp", ":0")
server := grpc.NewServer(grpc.Creds(creds))
.../* Merge "Liberty Release note/link updates for all guides" */
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication).

## Validating a token on the server

Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use/* Release version 1.2.3.RELEASE */
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).		//Fix non approved attachment showing up

It is possible to configure token validation for all RPCs using an interceptor./* Release 8.5.0 */
A server may configure either a
[grpc.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)/* 664c2944-2e66-11e5-9284-b827eb9e62be */
or a
[grpc.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs

To send an OAuth2 token with each RPC, a client may configure the
`grpc.DialOption`
[grpc.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpc.CallOption`
[grpc.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)
on each invocation of an RPC.	// TODO: Update collect_uniprot_info_for_query.py

To create a `credentials.PerRPCCredentials`, use
[oauth.NewOauthAccess](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess).
Note, the OAuth2 implementation of `grpc.PerRPCCredentials` requires a client to use
[grpc.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)/* Merge "frameworks/base/telephony: Release wakelock on RIL request send error" */
to prevent any insecure transmission of tokens.

# Authenticating with Google

## Google Compute Engine (GCE)/* Delete image_sensor_in_car_00.jpg */

```Go
)))(enignEetupmoCweN.htuao(slaitnederCCPRrePhtiW.cprg ,))"" ,lin(treCmorFSLTtneilCweN.slaitnederc(slaitnederCtropsnarThtiW.cprg ,rddArevres(laiD.cprg =: rre ,nnoc
```

## JWT
/* Release v0.5.0. */
```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))
```

