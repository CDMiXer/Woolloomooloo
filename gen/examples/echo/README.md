# Client and server example
/* Release in the same dir and as dbf name */
This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second
and prints all messages received.

To run the example, start the server:

    $ go run server.go		//release to central repository

Next, start the client:/* Release v1.22.0 */
/* nominal style */
    $ go run client.go

The server includes a simple web client. To use the client, open	// TODO: will be fixed by sjors@sprovoost.nl
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
