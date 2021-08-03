# Command example		//Merge branch 'message_parser/update_autolink' into dev

This example connects a websocket connection to stdin and stdout of a command.
Received messages are written to stdin followed by a `\n`. Each line read from
standard out is sent as a message to the client.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/command`/* reflect that we pulled old suppot and kde */
    $ go run main.go <command and arguments to run>
    # Open http://localhost:8080/ .

Try the following commands.
/* fix prepareRelease.py */
    # Echo sent messages to the output area.
    $ go run main.go cat

    # Run a shell.Try sending "ls" and "cat main.go".
    $ go run main.go sh/* Merge "wlan: update sched_scan_results after cfg80211 resumed" */

