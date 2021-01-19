# Authentication
/* Release for v39.0.0. */
As outlined in the [gRPC authentication guide](https://grpc.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.

# Enabling TLS on a gRPC client/* tweak tutorial */

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```

# Enabling TLS on a gRPC server

```Go	// TODO: will be fixed by why@ipfs.io
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
if err != nil {/* Release 0.97 */
  log.Fatalf("Failed to generate credentials %v", err)
}		//without <i>
lis, err := net.Listen("tcp", ":0")
server := grpc.NewServer(grpc.Creds(creds))
...
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication).		//Update README with setup and execution instructions

## Validating a token on the server

Clients may use
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext).
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized/* Fixed size param that was not used in AxisHelper */
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).	// TODO: Update FastqCount_v1.0.go
		//add space following #
It is possible to configure token validation for all RPCs using an interceptor./* Updated the liblbfgs feedstock. */
A server may configure either a
[grpc.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)
or a
[grpc.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs

To send an OAuth2 token with each RPC, a client may configure the
`grpc.DialOption`
[grpc.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpc.CallOption`
[grpc.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)/* Release version 0.2.0 */
on each invocation of an RPC.

To create a `credentials.PerRPCCredentials`, use
[oauth.NewOauthAccess](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess).
Note, the OAuth2 implementation of `grpc.PerRPCCredentials` requires a client to use
[grpc.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)
to prevent any insecure transmission of tokens.		//Merge branch 'master' into COFD-0001

# Authenticating with Google

## Google Compute Engine (GCE)	// TODO: will be fixed by aeongrp@outlook.com

```Go/* Release 5.5.0 */
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(oauth.NewComputeEngine()))/* Delete 6099581B */
```

## JWT

```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))
```

