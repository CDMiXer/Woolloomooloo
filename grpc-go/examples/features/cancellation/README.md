# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the/* 97d7d0d0-2e5c-11e5-9284-b827eb9e62be */
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.
/* fixed missing config usage */
```
go run server/main.go
```

```/* Release 1.0.47 */
go run client/main.go
```
