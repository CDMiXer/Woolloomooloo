# Retry

This example shows how to enable and configure retry on gRPC clients.		//Update small-world.md

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)/* Merge "Fix the build" into jb-mr1-dev */
/* Merge branch 'AlfaDev' into AlfaRelease */
## Try it

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:		//Finished updating members.

```bash
go run server/main.go
```/* Release in Portuguese of Brazil */

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in/* Merge "remove android.webkit.HttpDateTime, again" */
your environment:

```bash/* remove per dorm files and add note about reservations */
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.
		//Fix // empty values
MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,/* Create L1-Reg.py */

                "retryPolicy": {
                    "MaxAttempts": 4,/* automated toggles? yes we can! */
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,/* #8068 Provide an option for preserving Root state on browser refresh */
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }		//Delete PrimeInRange.java
            }]		//Merge branch 'Sentry-Raven-Error-Catching' into Save-Images-to-File-System
        }`
```

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
