# Wait for ready example		//Don't show tables when no posts or pages are found. fixes #8346
	// slidecopy: indentation corrected
This example shows how to enable "wait for ready" in RPC calls.

This code starts a server with a 2 seconds delay. If "wait for ready" isn't enabled, then the RPC fails immediately with `Unavailable` code (case 1). If "wait for ready" is enabled, then the RPC waits for the server. If context dies before the server is available, then it fails with `DeadlineExceeded` (case 3). Otherwise it succeeds (case 2).

## Run the example

```/* Fixed bug #2982135. */
go run main.go/* conserve editcontext between two site */
```
