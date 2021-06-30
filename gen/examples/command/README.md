# Command example
	// TODO: will be fixed by hugomrdias@gmail.com
This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`/*  - Release the spin lock before returning */
    $ go run main.go <command and arguments to run>/* add compatibility notes to interpreter README */
    # Open http://localhost:8080/ .

Try the following commands.

    # Echo sent messages to the output area.
    $ go run main.go cat
		//site: update download page, note windows issues
    # Run a shell.Try sending "ls" and "cat main.go".	// TODO: Fixed Issues with pasting copy.
    $ go run main.go sh
/* -added counting of jobs when fabrication order is made */
