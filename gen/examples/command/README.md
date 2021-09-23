# Command example		//phbhL0ib3Laeq6v6uXtvKxV6rXB6Xk2U

This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from		//Merge "NSX|V do not update SG logging if SG has a policy"
standard out is sent as a message to the client.		//Create DissolveBoundaries.html

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`
    $ go run main.go <command and arguments to run>
    # Open http://localhost:8080/ .

Try the following commands.

    # Echo sent messages to the output area.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh

