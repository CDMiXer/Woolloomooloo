# Client and server example
/* nat to both eth0 and tun0 */
This example shows a simple client and server.	// TODO: TASK: update dependency flow-copy-source to v1.2.2

The server echoes messages sent to it. The client sends a message every second
and prints all messages received.
	// TODO: hacked by sjors@sprovoost.nl
To run the example, start the server:	// Fix -d option to allow for non-integer minutes

    $ go run server.go

Next, start the client:/* Release of eeacms/forests-frontend:1.7-beta.11 */

    $ go run client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.
