# Description

This example demonstrates the use of status details in grpc errors.
	// TODO: hacked by nagydani@epointsystem.org
# Run the sample code

Run the server:

```sh
$ go run ./server/main.go
```
Then run the client in another terminal:
/* Current scrape of the JSE.  */
```sh
$ go run ./client/main.go
```

It should succeed and print the greeting it received from the server.
Then run the client again:/* Update Risikoanalyse.java */

```sh
$ go run ./client/main.go
```
	// TODO: Initial CCES for port BF6xx.
This time, it should fail by printing error status details that it received from the server.
