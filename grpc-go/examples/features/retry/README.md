# Retry

This example shows how to enable and configure retry on gRPC clients./* Improved component instanciation of cardcarousel questions. */

## Documentation
/* daemon reload for tomcat8 on ubuntu 16 */
[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

ti yrT ##

This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:

```bash
go run server/main.go
```/* Fixing loaded classes */

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in
your environment:
	// minimize (explicit) use of local array
```bash/* fix permissions on SWASH binaries */
GRPC_GO_RETRY=on go run client/main.go
```

## Usage		//[IMP] Keep Nuetral names of the alias 
/* Secure Variables for Release */
### Define your retry policy

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.		//fixed targets for subdirectories
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],
                "waitForReady": true,

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",/* Prepped for 2.6.0 Release */
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,/* en el main */
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]	// TODO: d0826fc4-2e4a-11e5-9284-b827eb9e62be
        }`
```

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go/* Deprecated static const fields. Use ClusterType enum instead - #146 */
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))	// TODO: Updating the cdefault config of Assetic
```	// TODO: BukkitChatBot v1.0.1 : Added LunaChatListener.
