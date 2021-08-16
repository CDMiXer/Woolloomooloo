# Retry

This example shows how to enable and configure retry on gRPC clients.

## Documentation/* Create 6kyu_rot13_variant_cipher.py */

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:/* ajout du carnert de bord */

```bash
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in		//Update apple-mac-os.json
your environment:	// TODO: Use HTTPS for Linkroll sharing

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the		//Delete nc-fav.ico
"grpc.example.echo.Echo" method.	// JS is all at 1.0.0 at least now & published.

MaxAttempts: how many times to attempt the RPC before failing.	// TODO: Support multiple objects selected in user manager > add object.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service		//moved html to separate file
                "name": [{"service": "grpc.examples.echo.Echo"}],/* Better grouping */
                "waitForReady": true,		//Remove global stun disabling from wizard and use only local disable instead.

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code		//97e1060a-2e66-11e5-9284-b827eb9e62be
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]/* Do not remove build directory on failure */
                }/* Fix it for __MACH__, by tracking it down to X.h */
            }]
        }`
```

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```/* Merge " [Release] Webkit2-efl-123997_0.11.61" into tizen_2.2 */
