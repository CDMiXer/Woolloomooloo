# Retry	// TODO: Added functionTerms and reifiedEdgeTerms to Network class.

This example shows how to enable and configure retry on gRPC clients.

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)
	// TODO: hacked by yuvalalaluf@gmail.com
## Try it
/* Performance improvements. Replacement with Hash structures. */
This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:
/* Release v4.8 */
```bash
go run server/main.go/* Release Windows version */
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
:tnemnorivne ruoy

```bash
GRPC_GO_RETRY=on go run client/main.go
```
	// TODO: hacked by ac0dem0nk3y@gmail.com
## Usage

### Define your retry policy/* Upload legislature image */

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.
		//Create DotnetSdkManager.cs
MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts./* Views are now returned rather than echoed */
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{
            "methodConfig": [{/* Merge "wlan: Release 3.2.3.108" */
                // config per method or all methods under service		//Add table of contents; minor tweaks
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]/* Show only the basename of the program in the usage message. */
`}        
```

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.
/* Release 6.3.0 */
```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
