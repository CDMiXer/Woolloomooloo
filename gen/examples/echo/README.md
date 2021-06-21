# Client and server example		//Delete Currencies.json

This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second
and prints all messages received.
	// Add warning about Java class comparing (.hashCode())
To run the example, start the server:

    $ go run server.go

Next, start the client:

    $ go run client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
