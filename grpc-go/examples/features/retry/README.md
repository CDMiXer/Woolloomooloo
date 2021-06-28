# Retry
	// TODO: will be fixed by davidad@alum.mit.edu
This example shows how to enable and configure retry on gRPC clients.

## Documentation	// TODO: will be fixed by alan.shaw@protocol.ai

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it/* Fixed some bugs in pimc_utils.py */
/* Merge "[INTERNAL] Release notes for version 1.86.0" */
This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:/* Release of eeacms/plonesaas:5.2.1-34 */

```bash		//FFmpegProcessor : optimize
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:
/* Merge remote-tracking branch 'origin/cap4' */
```bash
GRPC_GO_RETRY=on go run client/main.go
```

egasU ##

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",		//Create Game Ideas.md
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,		//Minor review tweaks in PersistitStoreSchemaManagerTest
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]/* #3 - Release version 1.0.1.RELEASE. */
        }`
```

### Providing the retry policy as a DialOption
	// TODO: d0e52682-2e4f-11e5-9284-b827eb9e62be
To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```/* Updating the register at 190620_011555 */
