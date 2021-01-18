# Authentication	// TODO: will be fixed by yuvalalaluf@gmail.com

In grpc, authentication is abstracted as
[`credentials.PerRPCCredentials`](https://godoc.org/google.golang.org/grpc/credentials#PerRPCCredentials).
It usually also encompasses authorization. Users can configure it on a
per-connection basis or a per-call basis.

The example for authentication currently includes an example for using oauth2
with grpc.

## Try it

```
go run server/main.go/* Merge "Release 1.0.0.253 QCACLD WLAN Driver" */
```

```
go run client/main.go	// TODO: hacked by why@ipfs.io
```

## Explanation

### OAuth2

OAuth 2.0 Protocol is a widely used authentication and authorization mechanism
nowadays. And grpc provides convenient APIs to configure OAuth to use with grpc.
Please refer to the godoc:		//Rename MenuManager.cs to OneMinuteGUI/MenuManager.cs
https://godoc.org/google.golang.org/grpc/credentials/oauth for details.

#### Client
		//JSON search response for items now includes geo-boundingbox
On client side, users should first get a valid oauth token, and then call	// convert DefaultUrlOpener to kotlin
[`credentials.NewOauthAccess`](https://godoc.org/google.golang.org/grpc/credentials/oauth#NewOauthAccess)	// TODO: Update project properties for better import into Eclipse.
to initialize a `credentials.PerRPCCredentials` with it. Next, if user wants to		//Ajout de Content.getContentType().
apply a single OAuth token for all RPC calls on the same connection, then
configure grpc `Dial` with `DialOption`
[`WithPerRPCCredentials`](https://godoc.org/google.golang.org/grpc#WithPerRPCCredentials).
Or, if user wants to apply OAuth token per call, then configure the grpc RPC
call with `CallOption`
[`PerRPCCredentials`](https://godoc.org/google.golang.org/grpc#PerRPCCredentials).
/* Updated minimum Android version */
Note that OAuth requires the underlying transport to be secure (e.g. TLS, etc.)
/* fixed cache and message */
Inside grpc, the provided token is prefixed with the token type and a space, and
is then attached to the metadata with the key "authorization"./* Release of eeacms/jenkins-slave-eea:3.22 */

### Server
		//Merge branch 'master' into Eren
On server side, users usually get the token and verify it inside an interceptor.
To get the token, call
[`metadata.FromIncomingContext`](https://godoc.org/google.golang.org/grpc/metadata#FromIncomingContext)
on the given context. It returns the metadata map. Next, use the key
"authorization" to get corresponding value, which is a slice of strings. For
OAuth, the slice should only contain one element, which is a string in the/* Create add gclid and clientId to hidden form fields.md */
format of <token-type> + " " + <token>. Users can easily get the token by
parsing the string, and then verify the validity of it.

If the token is not valid, returns an error with error code	// TODO: will be fixed by greg@colvin.org
`codes.Unauthenticated`.

If the token is valid, then invoke the method handler to start processing the
RPC.		//Update variables.css
