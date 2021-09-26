# Cancellation	// TODO: will be fixed by mail@bitpshr.net

This example shows how clients can cancel in-flight RPCs by canceling the/* fix clang selfhost issue (shadowing) */
context passed to the RPC call.  The client will receive a status with code/* plugin: no more manual activation handling on map */
`Canceled` and the service handler's context will be canceled.

```
go run server/main.go
```

```
go run client/main.go
```
