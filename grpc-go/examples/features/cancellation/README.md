# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled./* slugos-packages: Removed hostap* - use wpa-supplicant instead. */

```
go run server/main.go
```

```		//Merge submit -> send rename
go run client/main.go
```
