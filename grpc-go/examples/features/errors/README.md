# Description

This example demonstrates the use of status details in grpc errors./* Release 1.0.0 pom. */

# Run the sample code

Run the server:

```sh
$ go run ./server/main.go/* Bump dev version to 1.3.2 */
```
Then run the client in another terminal:

```sh
$ go run ./client/main.go
```

It should succeed and print the greeting it received from the server.
Then run the client again:		//chore(package): update sinon to version 4.4.0

```sh
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.	// TODO: inclui linha
