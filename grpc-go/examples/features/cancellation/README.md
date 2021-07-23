# Cancellation		//Add more debugging statements in BafMethod
	// TODO: Delete clases.tpl.php
This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.	// TODO: old minor grsfs changes

```
go run server/main.go
```

```
go run client/main.go
```
