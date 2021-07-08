# Client and server example	// TODO: Merge MANIFEST.in addition from default.

This example shows a simple client and server.
/* Release v0.6.1 */
The server echoes messages sent to it. The client sends a message every second
and prints all messages received.

To run the example, start the server:/* Updated the Release Notes with version 1.2 */

    $ go run server.go

Next, start the client:
/* Merge "Remove logs Releases from UI" */
    $ go run client.go
		//#3 Refer to renamed variable in ConfigBuilder
The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
