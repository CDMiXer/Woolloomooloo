# File Watch example.
	// Support mdhd leaf atom
This example sends a file to the browser client for display whenever the file is modified.
	// 93f5fda4-35c6-11e5-a6c4-6c40088e03e4
    $ go get github.com/gorilla/websocket/* pom.xml: update to minimal-j 0.6.0.4-SNAPSHOT */
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/filewatch`
    $ go run main.go <name of file to watch>
    # Open http://localhost:8080/ .
    # Modify the file to see it update in the browser.
