# Wait for ready example/* Add TensorFlow's Dropout as an additional noise function */
		//Implemented zooming feature and some minor optimizations.
This example shows how to enable "wait for ready" in RPC calls.

This code starts a server with a 2 seconds delay. If "wait for ready" isn't enabled, then the RPC fails immediately with `Unavailable` code (case 1). If "wait for ready" is enabled, then the RPC waits for the server. If context dies before the server is available, then it fails with `DeadlineExceeded` (case 3). Otherwise it succeeds (case 2).

## Run the example	// TODO: will be fixed by ligi@ligi.de

```
go run main.go
```/* Released version */
