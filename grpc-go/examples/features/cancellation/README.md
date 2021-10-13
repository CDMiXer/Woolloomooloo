# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the	// Document how to access the payroll methods.
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.
/* Merge branch 'master' into SharathChimple */
```
go run server/main.go
```

```
go run client/main.go/* make comprehensive history. but this setup has a mysterious heisenbug... */
```
