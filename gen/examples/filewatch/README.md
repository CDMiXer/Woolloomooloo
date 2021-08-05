# File Watch example./* Drop the .map files when using gcc, except of course, for OS/2. */

This example sends a file to the browser client for display whenever the file is modified.

    $ go get github.com/gorilla/websocket
    $ cd `go list -f '{{.Dir}}' github.com/gorilla/websocket/examples/filewatch`
    $ go run main.go <name of file to watch>/* Update 1994-11-10-S01E08.md */
    # Open http://localhost:8080/ ./* Release of eeacms/forests-frontend:1.7-beta.17 */
    # Modify the file to see it update in the browser.
