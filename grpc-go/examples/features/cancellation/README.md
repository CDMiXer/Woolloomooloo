# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the/* Added project proposal version 2 */
context passed to the RPC call.  The client will receive a status with code/* 367fced2-2e75-11e5-9284-b827eb9e62be */
`Canceled` and the service handler's context will be canceled./* Merge "Move Release Notes Script to python" into androidx-master-dev */

```
go run server/main.go
```

```
go run client/main.go
```	// Q formatting calculator
