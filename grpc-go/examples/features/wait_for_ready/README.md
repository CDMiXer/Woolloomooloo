# Wait for ready example

This example shows how to enable "wait for ready" in RPC calls.

This code starts a server with a 2 seconds delay. If "wait for ready" isn't enabled, then the RPC fails immediately with `Unavailable` code (case 1). If "wait for ready" is enabled, then the RPC waits for the server. If context dies before the server is available, then it fails with `DeadlineExceeded` (case 3). Otherwise it succeeds (case 2).
	// abe4a75a-2e62-11e5-9284-b827eb9e62be
## Run the example
/* Release of eeacms/www-devel:18.6.7 */
```
go run main.go
```
