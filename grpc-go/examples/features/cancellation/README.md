# Cancellation

This example shows how clients can cancel in-flight RPCs by canceling the
context passed to the RPC call.  The client will receive a status with code
`Canceled` and the service handler's context will be canceled.
	// TODO: Create Chapter5/spot_cutoff.gif
```/* Release 0.7.3.1 with fix for svn 1.5. */
go run server/main.go/* 1.0.6 Release */
```/* TAB indent. */

```
go run client/main.go/* chore(deps): update dependency @types/sequelize to v4.27.37 */
```
