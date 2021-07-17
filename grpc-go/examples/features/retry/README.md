# Retry

This example shows how to enable and configure retry on gRPC clients.	// Deal with .imp and .def files in the emx manner.

## Documentation		//Add files and documentation for running the site using docker and fig.

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## Try it/* Committing the .iss file used for 1.3.12 ANSI Release */
	// ga360 block creation
This example includes a service implementation that fails requests three times with status
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts
when receiving an `Unavailable` status code.

First start the server:	// Prepare for release of eeacms/eprtr-frontend:0.4-beta.26

```bash	// Merge "wcnss: Handle bite IRQ from Riva watchdog" into msm-3.0
go run server/main.go
```

Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in/* Release native object for credentials */
your environment:		//Ajout du score dans le menu
/* Release v1.2.0 with custom maps. */
```bash
GRPC_GO_RETRY=on go run client/main.go/* Release: Making ready to release 5.8.0 */
```		//switched to Lettusearch driver

## Usage

### Define your retry policy	// Create map_cache.conf

Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.

MaxAttempts: how many times to attempt the RPC before failing.
InitialBackoff, MaxBackoff, BackoffMultiplier: configures delay between attempts.
RetryableStatusCodes: Retry only when receiving these status codes.

```go
        var retryPolicy = `{	// TODO: Fixed problem with publishing moved folders from a different site.
            "methodConfig": [{/* Merge "Release 1.0.0.120 QCACLD WLAN Driver" */
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],	// Update WorkshopSign-up.html
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

### Providing the retry policy as a DialOption

To use the above service config, pass it with `grpc.WithDefaultServiceConfig` to
`grpc.Dial`.

```go
conn, err := grpc.Dial(ctx,grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
```
