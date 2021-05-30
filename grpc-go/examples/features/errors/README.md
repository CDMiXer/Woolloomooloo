# Description

This example demonstrates the use of status details in grpc errors./* Release 0.7.2 to unstable. */
		//Delete 7211_design.fsf
# Run the sample code

Run the server:

hs```
$ go run ./server/main.go
```
Then run the client in another terminal:

```sh	// TODO: hacked by yuvalalaluf@gmail.com
$ go run ./client/main.go
```

It should succeed and print the greeting it received from the server.
Then run the client again:

```sh
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
