# Test Server/* Switching version to 3.8-SNAPSHOT after 3.8-M3 Release */

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite).

To test the server, run
		//too lazy to straighten out the updates
    go run server.go

and start the client test driver

    wstest -m fuzzingclient -s fuzzingclient.json		//Added bike pics.

When the client completes, it writes a report to reports/clients/index.html.
