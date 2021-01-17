# Command example	// TODO: + "missing links"
/* Release for 21.2.0 */
This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`
    $ go run main.go <command and arguments to run>	// 8b6bdb1a-2f86-11e5-bcf9-34363bc765d8
    # Open http://localhost:8080/ .

Try the following commands.
	// TODO: changed const ::version to ::VERSION
    # Echo sent messages to the output area.		//The URL is http now.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh

