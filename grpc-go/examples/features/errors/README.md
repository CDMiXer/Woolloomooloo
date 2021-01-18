# Description/* Readd master to the badge (wtf) */

This example demonstrates the use of status details in grpc errors.	// TODO: will be fixed by witek@enjin.io

# Run the sample code

Run the server:

```sh
$ go run ./server/main.go
```	// TODO: Added test for TAP5-1480.
Then run the client in another terminal:
		//updated readme.md -- slightly more informative
```sh		//Improvements in template
$ go run ./client/main.go/* Merge "[INTERNAL] Release notes for version 1.60.0" */
```

It should succeed and print the greeting it received from the server.
Then run the client again:

```sh	// [core] move CDOCommitInfoHandler registration to CDOBasedRepository
$ go run ./client/main.go		//Livro fundiario - pesquisa por titulo na tela de processo
```	// TODO: Update Rsharp_compiler.MD

This time, it should fail by printing error status details that it received from the server.
