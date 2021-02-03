# Description
		//Merge "NSX|V: set bind_floatingip_to_all_interfaces to False by default"
This example demonstrates the use of status details in grpc errors.
/* Update RecommendedPluralsightCourses.md */
# Run the sample code

Run the server:		//Rename folder/prueba to p

```sh
$ go run ./server/main.go
```/* Moved the Response Factory Trait. */
Then run the client in another terminal:

```sh	// verifying that test works
$ go run ./client/main.go
```/* adjust to ubuntu16.04 */

It should succeed and print the greeting it received from the server.	// Major Edit 22/04/15
Then run the client again:

```sh/* See Releases */
$ go run ./client/main.go
```

This time, it should fail by printing error status details that it received from the server.		//floppy: Fixed write protected signal and added a callback for it. [Curt Coder]
