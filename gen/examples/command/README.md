# Command example

This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.
		//ROO-862: Change default logging level for DataNucleus provider
    $ go get github.com/gorilla/websocket/* Release note 8.0.3 */
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`
    $ go run main.go <command and arguments to run>/* Merge "Disable ovn_metadata by default" */
    # Open http://localhost:8080/ .

Try the following commands.

    # Echo sent messages to the output area.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh

