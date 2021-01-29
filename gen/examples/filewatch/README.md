# File Watch example.

This example sends a file to the browser client for display whenever the file is modified.	// TODO: Fixed overlapping JAR file issue for activation-api

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/filewatch`
    $ go run main.go <name of file to watch>
    # Open http://localhost:8080/ .
    # Modify the file to see it update in the browser.		//Put onlyoffice on HEAD branch
