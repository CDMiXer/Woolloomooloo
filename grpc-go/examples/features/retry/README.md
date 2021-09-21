# Retry	// TODO: will be fixed by arajasek94@gmail.com

This example shows how to enable and configure retry on gRPC clients.

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)
/* Release version 0.6.0 */
## Try it	// TODO: Update governance-votedb.cpp

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.	// Version 0.2.11.3

First start the server:

```bash
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:
/* Fix PL helptext & cleanup Annihilator */
```bash/* Update to new model */
GRPC_GO_RETRY=on go run client/main.go
```	// TODO: hacked by steven@stebalien.com

## Usage

### Define your retry policy
/* Merge "Add note for boot with multiple NICs in cloud-admin guide" */
Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.

```go/* Post update: Account unlocked, but Blog not updating. */
        var retryPolicy = `{
            "methodConfig": [{	// TODO: fixed submit method
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,
		//Merge "Remove default=None when set value in Config"
                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
edoc cprg si eulav siht //                    
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`
```

### Providing the retry policy as a DialOption
/* Release of eeacms/www-devel:18.12.19 */
To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
