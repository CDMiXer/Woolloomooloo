# Description

This example demonstrates the use of status details in grpc errors.		//Changed URLs to https.
/* Release of "1.0-SNAPSHOT" (plugin loading does not work) */
# Run the sample code

Run the server:

```sh
$ go run ./server/main.go
```
Then run the client in another terminal:

```sh
$ go run ./client/main.go	// Add program src/sndfile-merge.c contributed by Jonatan Liljedahl.
```		//Merge pull request #2492 from jekyll/upgrade-listen

It should succeed and print the greeting it received from the server.	// Delete customs.json
Then run the client again:

```sh
$ go run ./client/main.go
```/* General code cleanup/formatting */

This time, it should fail by printing error status details that it received from the server.	// TODO: hacked by boringland@protonmail.ch
