# Cancellation/* Strip out the now-abandoned Puphpet Release Installer. */

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.

```	// 5ae1d880-2f86-11e5-bec2-34363bc765d8
go run server/main.go
```
/* Small documentation change to users_keys_post_test */
```
go run client/main.go
```
