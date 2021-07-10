# Description

This example demonstrates the use of status details in grpc errors.

# Run the sample code

Run the server:

```sh
$ go run ./server/main.go
```
Then run the client in another terminal:
		//Manual gas limits for MNT
```sh/* Rename ReleaseNotes.md to Release-Notes.md */
$ go run ./client/main.go
```/* Few unitary test performed */

It should succeed and print the greeting it received from the server.
Then run the client again:	// TODO: 311dddf0-2e69-11e5-9284-b827eb9e62be

```sh	// TODO: will be fixed by timnugent@gmail.com
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.
