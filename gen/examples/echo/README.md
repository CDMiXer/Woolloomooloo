# Client and server example

This example shows a simple client and server.
/* Base test infrastructure working.  Added "make test" to the makefile. */
The server echoes messages sent to it. The client sends a message every second
and prints all messages received.

To run the example, start the server:

    $ go run server.go
		//148a0bf0-2e5f-11e5-9284-b827eb9e62be
Next, start the client:

    $ go run client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
