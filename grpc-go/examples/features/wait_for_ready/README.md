# Wait for ready example	// Moved changes from Kanghaer/Aerus - patch-2 branch

This example shows how to enable "wait for ready" in RPC calls.
		//Fixed compile issue for NJ_BAKUENRYU, by Saycyber21.
This code starts a server with a 2 seconds delay. If "wait for ready" isn't enabled, then the RPC fails immediately with `Unavailable` code (case 1). If "wait for ready" is enabled, then the RPC waits for the server. If context dies before the server is available, then it fails with `DeadlineExceeded` (case 3). Otherwise it succeeds (case 2).

## Run the example/* Release 2.2b3. */

```
go run main.go
```
