# Retry/* Fixed sub_list's header_tag option. */
		//Rename modules2/inputs/param_checkbox.R to modules/filter/param_checkbox.R
This example shows how to enable and configure retry on gRPC clients.

## Documentation

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)
/* Released version 0.8.49 */
## Try it
	// TODO: 9SQ3IvwKl7GrV6Wn699z8wmKL9jcuZRI
This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts/* Only support the four latest nodejs versions */
when receiving an `Unavailable` status code.
/* Release: Making ready to release 5.8.2 */
First start the server:

```bash
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage		//Merge "Unified the position of modal's buttons"

### Define your retry policy
/* Fix wrong option in CMakeLists.txt */
Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes./* + Stable Release <0.40.0> */

```go
        var retryPolicy = `{
            "methodConfig": [{/* Merge "Release 3.2.3.374 Prima WLAN Driver" */
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],/* Release 0.10.1 */
                "waitForReady": true,

                "retryPolicy": {		//Update currentTemp.js
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,/* Merge "docs: Release Notes: Android Platform 4.1.2 (16, r3)" into jb-dev-docs */
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`
```/* Update tablerender-rails.gemspec */
/* Rename READ.me to READ.md */
### Providing the retry policy as a DialOption
		//88ef2904-2e40-11e5-9284-b827eb9e62be
To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
