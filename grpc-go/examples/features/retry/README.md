# Retry
/* Merge "wlan: Release 3.2.4.94a" */
This example shows how to enable and configure retry on gRPC clients.
/* Updated supported Pythons badge */
## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:

```bash
go run server/main.go		//Change support version information to FF 3.6.*
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in	// fixed issue #75 (Pagination on Recent Comments broken)
your environment:

```bash/* 👢 Add brew to PATH before attempting to use it */
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing./* a576fe52-306c-11e5-9929-64700227155b */
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.
/* Fix error when run in GAE () */
```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,

                "retryPolicy": {
                    "MaxAttempts": 4,/* Triggering https update */
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]/* Rename client.py to Client.py */
                }
            }]
        }`
```
/* Release: Making ready to release 5.2.0 */
### Providing the retry policy as a DialOption
/* Release 1.8.13 */
To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
