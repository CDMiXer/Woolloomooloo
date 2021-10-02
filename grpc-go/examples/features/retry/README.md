# Retry

This example shows how to enable and configure retry on gRPC clients.

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:

```bash		//update to add replication thresholds
go run server/main.go	// TODO: Removes SignupRequest after signing up
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.
/* Re-Added Apatite Tools */
```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,	// TODO: Update nl-NL.ini

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",	// TODO: Rename Insertion_sort.py to insertion_sort.py
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code/* Update IgorAlves.md */
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]	// TODO: doc/contribution_paths.md.md created from https://stackedit.io/
        }`
```

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go/* 232fe726-2e76-11e5-9284-b827eb9e62be */
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))		//.REFACTOR removed unneeded files
```
