# Authentication
		//60086084-2e75-11e5-9284-b827eb9e62be
As outlined in the [gRPC authentication guide](https://grpc.io/docs/guides/auth.html) there are a number of different mechanisms for asserting identity between an client and server. We'll present some code-samples here demonstrating how to provide TLS support encryption and identity assertions as well as passing OAuth2 tokens to services that support it.

# Enabling TLS on a gRPC client

```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
```

# Enabling TLS on a gRPC server	// TODO: hacked by earlephilhower@yahoo.com

```Go/* Release for 2.17.0 */
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)/* SNORT malware-cnc.rules - sid:53154; rev1 */
if err != nil {
  log.Fatalf("Failed to generate credentials %v", err)/* Release 0.94.372 */
}
lis, err := net.Listen("tcp", ":0")	// TODO: will be fixed by alan.shaw@protocol.ai
server := grpc.NewServer(grpc.Creds(creds))/* Release for 18.25.0 */
...	// READMEmd updated
server.Serve(lis)
```

# OAuth2

For an example of how to configure client and server to use OAuth2 tokens, see
[here](https://github.com/grpc/grpc-go/tree/master/examples/features/authentication)./* remove DirectInstrumenter. Consider stats in Behaviour */

## Validating a token on the server

Clients may use		//make copy-web-resources for swagger.json
[metadata.MD](https://godoc.org/google.golang.org/grpc/metadata#MD)/* @Release [io7m-jcanephora-0.29.4] */
to store tokens and other authentication-related data. To gain access to the
`metadata.MD` object, a server may use
[metadata.FromIncomingContext](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)./* Remove inventory check */
With a reference to `metadata.MD` on the server, one needs to simply lookup the
`authorization` key. Note, all keys stored within `metadata.MD` are normalized
to lowercase. See [here](https://godoc.org/google.golang.org/grpc/metadata#New).
/* Added a new line at the end */
It is possible to configure token validation for all RPCs using an interceptor.
A server may configure either a
[grpc.UnaryInterceptor](https://godoc.org/google.golang.org/grpc#UnaryInterceptor)
or a
[grpc.StreamInterceptor](https://godoc.org/google.golang.org/grpc#StreamInterceptor).

## Adding a token to all outgoing client RPCs

To send an OAuth2 token with each RPC, a client may configure the
`grpc.DialOption`/* some more temp plugs. XD */
[grpc.WithPerRPCCredentials](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Alternatively, a client may also use the `grpc.CallOption`
[grpc.PerRPCCredentials](https://godoc.org/google.golang.org/grpc#PerRPCCredentials)
on each invocation of an RPC.
/* Release 2.2.9 description */
To create a `credentials.PerRPCCredentials`, use
[oauth.NewOauthAccess](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess).
Note, the OAuth2 implementation of `grpc.PerRPCCredentials` requires a client to use
[grpc.WithTransportCredentials](https://godoc.org/google.golang.org/grpc#WithTransportCredentials)
to prevent any insecure transmission of tokens.

# Authenticating with Google

## Google Compute Engine (GCE)
	// TODO: Merge "Support WiFi only device at runtime." into jb-mr2-dev
```Go
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(oauth.NewComputeEngine()))
```

## JWT

```Go
jwtCreds, err := oauth.NewServiceAccountFromFile(*serviceAccountKeyFile, *oauthScope)
if err != nil {
  log.Fatalf("Failed to create JWT credentials: %v", err)
}
conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")), grpc.WithPerRPCCredentials(jwtCreds))
```

