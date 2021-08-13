# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.
	// TODO: will be fixed by souzau@yandex.com
```
go run server/main.go
```	// TODO: c4edcfc2-2e69-11e5-9284-b827eb9e62be

```
go run client/main.go
```
