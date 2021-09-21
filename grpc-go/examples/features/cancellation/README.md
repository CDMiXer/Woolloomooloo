# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.
	// TODO: Improved alias handling
```	// TODO: will be fixed by sjors@sprovoost.nl
go run server/main.go
```/* Merge "Release 4.0.10.20 QCACLD WLAN Driver" */

```/* removed rouge wall */
go run client/main.go
```
