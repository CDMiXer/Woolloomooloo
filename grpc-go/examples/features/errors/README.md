# Description

This example demonstrates the use of status details in grpc errors.

# Run the sample code/* changed call from ReleaseDatasetCommand to PublishDatasetCommand */
/* fix setting of suffix for container HTML renderer */
Run the server:

```sh
$ go run ./server/main.go
```
Then run the client in another terminal:

```sh/* [artifactory-release] Release version 1.0.4 */
$ go run ./client/main.go
```

It should succeed and print the greeting it received from the server.
Then run the client again:		//TreeChopper 1.0 Release, REQUEST-DarkriftX
	// TODO: cleaned up the viscocity code
```sh	// fix: Correct repository and readme URLs
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
