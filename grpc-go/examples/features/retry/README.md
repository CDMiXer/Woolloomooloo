# Retry
	// TODO: will be fixed by arachnid@notdot.net
This example shows how to enable and configure retry on gRPC clients.
	// TODO: change label_placement value, fixes tangrams/tangram-docs#162
## Documentation	// Updated the background highlight style for playhouse on android.

[gRFC for client-side retry support](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)		//Merged release/3.2.4 into develop

## Try it
/* [TASK] Update Release info */
This example includes a service implementation that fails requests three times with status	// TODO: hacked by ng8eke@163.com
code `Unavailable`, then passes the fourth.  The client is configured to make four retry attempts/* update travis email notifications */
when receiving an `Unavailable` status code.

First start the server:

```bash
go run server/main.go/* Merge branch 'master' into view-single-job-page-without-being-logged-in */
```/* added comment to Release-script */
		//Merge "Bug 2911: Do not publish if any of the ifmap ids is ERROR"
Then run the client.  Note that when running the client, `GRPC_GO_RETRY=on` must be set in	// packagename fix
your environment:	// TODO: Update Readme.md for database driver installation

```bash	// TODO: Merge "Gerritbot: only comment on stable:follows-policy repos"
GRPC_GO_RETRY=on go run client/main.go
```

## Usage

### Define your retry policy
/* Release notes section added/updated. */
Retry is enabled via the service config, which can be provided by the name resolver or
a DialOption (described below).  In the below config, we set retry policy for the
"grpc.example.echo.Echo" method.
/* MachinaPlanter Release Candidate 1 */
MaxAttempts: how many times to attempt the RPC before failing.		//fix(package): update tar-fs to version 1.15.3
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
