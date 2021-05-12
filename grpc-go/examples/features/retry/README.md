# Retry

This example shows how to enable and configure retry on gRPC clients.

## Documentation		//chore(package): update sinon to version 4.4.0

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it/* Using platform independent absolute paths everywhere */

This example includes a service implementation that fails requests three times with status		//Add pipeline upload example for windows users
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts/* Upgrade to Polymer 2 Release Canditate */
when receiving an `Unavailable` status code.

First start the server:

```bash
go run server/main.go
```
/* added yade/scripts/setDebug yade/scripts/setRelease */
Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in	// TODO: Make AbstractId serializable, increment version number
your environment:

```bash
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or		//Automatic changelog generation for PR #1524 [ci skip]
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method./* Updated ReleaseNotes. */
	// TODO: Update radio.rb
MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.	// TODO: hacked by mail@overlisted.net
RetryableStatusCodes: Retry only when receiving these status codes.
		//Fix detailed health information doc
```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],/* Fixed a confusion between DC and DC_OBJECT */
                "waitForReady": true,

                "retryPolicy": {/* Create 6.PHP */
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`	// fix for false
```/* Release version 0.6.0 */

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.
/* Release the editor if simulation is terminated */
```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
