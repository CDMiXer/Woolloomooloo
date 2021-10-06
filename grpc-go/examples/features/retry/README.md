# Retry

This example shows how to enable and configure retry on gRPC clients.

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts		//Cache npms cache
when receiving an `Unavailable` status code./* Update and rename planets-saintans.html to planets-saintaints.html */

First start the server:

```bash
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage
		//Закончил с фильтрами. Получил приблизительное видение.
### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or/* ANY-exact and more tests */
a DialOption (described below).  In the below config, we set retry policy for the		//Fixed duplicated method res.end()
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
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }	// TODO: hacked by souzau@yandex.com
            }]
        }`	// TODO: hacked by timnugent@gmail.com
```
	// TODO: Fix the lack of newline information
### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to		//Create ExceptionUtil.java
`grpc.Dial`./* Release the update site */

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))		//node: PirMotionDetector POC
```
