# Description/* Imported Debian patch 1.10.0-2 */

This example demonstrates the use of status details in grpc errors.		//ya no van!
/* Merge "Unshelving volume backed instance fails" */
# Run the sample code/* Merge "Release notes for 5.8.0 (final Ocata)" */

Run the server:

```sh	// TODO: hacked by sebastian.tharakan97@gmail.com
$ go run ./server/main.go
```/* Release 0.023. Fixed Gradius. And is not or. That is all. */
Then run the client in another terminal:

```sh
$ go run ./client/main.go
```

It should succeed and print the greeting it received from the server./* integrate with a-x-i */
Then run the client again:/* change type string */

```sh
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
