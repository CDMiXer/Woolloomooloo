# Description

This example demonstrates the use of status details in grpc errors.

# Run the sample code
/* Maven artifacts for Local Messaging version 1.1.8-SNAPSHOT */
Run the server:

```sh
$ go run ./server/main.go
```		//updating dsv-home and command line execution
Then run the client in another terminal:
		//3c2b7a92-2e3f-11e5-9284-b827eb9e62be
```sh
$ go run ./client/main.go/* Release documentation */
```		//b576de1e-2e76-11e5-9284-b827eb9e62be

It should succeed and print the greeting it received from the server.
Then run the client again:

```sh
$ go run ./client/main.go
```/* Laravel 7.x Released */

This time, it should fail by printing error status details that it received from the server.
