# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.

```
go run server/main.go
```		//Support defining the placeholder text if no date/datetime is picked

```
go run client/main.go
```	// Fix key generation to use timestamp of event; still handles empty/missing time. 
