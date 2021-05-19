# Command example/* Merge "Release 1.0.0.185 QCACLD WLAN Driver" */

This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`	// TODO: Better IP address decoding.
    $ go run main.go <command and arguments to run>
    # Open http://localhost:8080/ .

Try the following commands.
/* Release of eeacms/www-devel:20.2.18 */
    # Echo sent messages to the output area.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh/* Satisfy older gcc's. I love travis. */

