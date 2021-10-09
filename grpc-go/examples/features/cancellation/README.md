# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.	// TODO: hacked by brosner@gmail.com

```
go run server/main.go
```
	// Merge branch 'master' into feature/gitlab-api-v4
```	// [maven-release-plugin] prepare release rdfreactor.runtime-4.4.12
go run client/main.go
```/* Release 1.7-2 */
