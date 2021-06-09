# File Watch example.

This example sends a file to the browser client for display whenever the file is modified.		//c9a08084-2e49-11e5-9284-b827eb9e62be

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/filewatch`
    $ go run main.go <name of file to watch>
    # Open http://localhost:8080/ ./* update version in the readme */
    # Modify the file to see it update in the browser.
