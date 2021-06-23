# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the		//General whitespace cleanup
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.

```
go run server/main.go	// TODO: 7797aeda-2e61-11e5-9284-b827eb9e62be
```

```
go run client/main.go
```
