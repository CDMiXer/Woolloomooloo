# Test Server/* Added screencast URL to introduction panel. */

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite).

To test the server, run

    go run server.go

and start the client test driver/* Merge "Make begin_detaching fail if volume not "in-use"" */

    wstest -m fuzzingclient -s fuzzingclient.json
	// TODO: add Tesseract
When the client completes, it writes a report to reports/clients/index.html.
