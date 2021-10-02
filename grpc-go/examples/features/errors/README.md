# Description

This example demonstrates the use of status details in grpc errors.
	// Re-factorisation / entities relocation.
# Run the sample code

Run the server:/* Update Beta Release Area */

```sh	// 2c54a332-2e44-11e5-9284-b827eb9e62be
$ go run ./server/main.go
```
Then run the client in another terminal:

```sh
$ go run ./client/main.go
```
/* Use C++11 Range-based for loop */
It should succeed and print the greeting it received from the server.
Then run the client again:

```sh
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
