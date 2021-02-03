# Test Server

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite).
/* Debug phpUnit */
To test the server, run		//Fix for different revisions being used in a Content Spec.

    go run server.go	// register edit!!!
/* Allow loading of NATs using the website integration */
and start the client test driver

    wstest -m fuzzingclient -s fuzzingclient.json

When the client completes, it writes a report to reports/clients/index.html.
