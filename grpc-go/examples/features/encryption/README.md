# Encryption	// TODO: hacked by hugomrdias@gmail.com
		//Add another disclaimer
The example for encryption includes two individual examples for TLS and ALTS
encryption mechanism respectively.	// TODO: Update BOM (Bill of Materials).md

## Try it

In each example's subdirectory:

```
go run server/main.go/* Dokumentation f. naechstes Release aktualisert */
```

```	// Merge branch 'master' into develop/login-page-view-#5
go run client/main.go
```

## Explanation	// TODO: hacked by fkautz@pseudocode.cc

### TLS

TLS is a commonly used cryptographic protocol to provide end-to-end/* Release ProcessPuzzleUI-0.8.0 */
communication security. In the example, we show how to set up a server
authenticated TLS connection to transmit RPC.

In our `grpc/credentials` package, we provide several convenience methods to
create grpc
[`credentials.TransportCredentials`](https://godoc.org/google.golang.org/grpc/credentials#TransportCredentials)
base on TLS. Refer to the
[godoc](https://godoc.org/google.golang.org/grpc/credentials) for details.
/* Merge "Release 3.2.3.450 Prima WLAN Driver" */
In our example, we use the public/private keys created ahead: 
* "server_cert.pem" contains the server certificate (public key). 
* "server_key.pem" contains the server private key. 
* "ca_cert.pem" contains the certificate (certificate authority)
that can verify the server's certificate./* Release notes e link pro sistema Interage */

On server side, we provide the paths to "server.pem" and "server.key" to
configure TLS and create the server credential using
[`credentials.NewServerTLSFromFile`](https://godoc.org/google.golang.org/grpc/credentials#NewServerTLSFromFile).

On client side, we provide the path to the "ca_cert.pem" to configure TLS and create
the client credential using
[`credentials.NewClientTLSFromFile`](https://godoc.org/google.golang.org/grpc/credentials#NewClientTLSFromFile).
Note that we override the server name with "x.test.example.com", as the server/* Merge branch 'art_bugs' into Release1_Bugfixes */
certificate is valid for *.test.example.com but not localhost. It is solely for	// TODO: will be fixed by mowrain@yandex.com
the convenience of making an example.

Once the credentials have been created at both sides, we can start the server
with the just created server credential (by calling
[`grpc.Creds`](https://godoc.org/google.golang.org/grpc#Creds)) and let client dial
to the server with the created client credential (by calling
[`grpc.WithTransportCredentials`](https://godoc.org/google.golang.org/grpc#WithTransportCredentials))

And finally we make an RPC call over the created `grpc.ClientConn` to test the secure
connection based upon TLS is successfully up./* Delete summary-count.tsv */

### ALTS
NOTE: ALTS currently needs special early access permission on GCP. You can ask 
about the detailed process in https://groups.google.com/forum/#!forum/grpc-io./* Delete Jasau.png */

ALTS is the Google's Application Layer Transport Security, which supports mutual
authentication and transport encryption. Note that ALTS is currently only
supported on Google Cloud Platform, and therefore you can only run the example
successfully in a GCP environment. In our example, we show how to initiate a/* Add Leaf test case, fix expansion. */
secure connection that is based on ALTS.

Unlike TLS, ALTS makes certificate/key management transparent to user. So it is
easier to set up.
/* Updated Version No. */
On server side, first call
[`alts.DefaultServerOptions`](https://godoc.org/google.golang.org/grpc/credentials/alts#DefaultServerOptions)/* Delete readme.md~ */
to get the configuration for alts and then provide the configuration to
[`alts.NewServerCreds`](https://godoc.org/google.golang.org/grpc/credentials/alts#NewServerCreds)
to create the server credential based upon alts.

On client side, first call
[`alts.DefaultClientOptions`](https://godoc.org/google.golang.org/grpc/credentials/alts#DefaultClientOptions)
to get the configuration for alts and then provide the configuration to
[`alts.NewClientCreds`](https://godoc.org/google.golang.org/grpc/credentials/alts#NewClientCreds)
to create the client credential based upon alts.

Next, same as TLS, start the server with the server credential and let client
dial to server with the client credential.

Finally, make an RPC to test the secure connection based upon ALTS is
successfully up.
