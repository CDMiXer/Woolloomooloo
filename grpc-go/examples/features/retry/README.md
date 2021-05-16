# Retry
/* Added pip3 to update function */
This example shows how to enable and configure retry on gRPC clients.	// TODO: new template version with comments

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)/* f7200c46-2e62-11e5-9284-b827eb9e62be */
		//New version of Silver, Blue &amp; Gold - 1.06
## Try it/* Merge branch 'vNext_yarong_417_812' into vNext_yarong_417_812_artemjackson_1190 */

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code./* Release REL_3_0_5 */

First start the server:/* Released DirectiveRecord v0.1.23 */

```bash
go run server/main.go
```
/* Release version 0.0.6 */
Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage
	// TODO: f5c8881a-2e4d-11e5-9284-b827eb9e62be
### Define your retry policy
/* Release 1.0.3b */
Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.		//Add note on updating translations.
RetryableStatusCodes: Retry only when receiving these status codes.
/* tup build: add STM32F1 chip sources to compilation */
```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service/* 5fadc8f2-2e63-11e5-9284-b827eb9e62be */
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
            }]
        }`
```
/* added jacoco / coveralls */
### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.
	// TODO: identity of viewpitch in software and gl
```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
