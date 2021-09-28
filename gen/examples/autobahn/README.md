# Test Server

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite)./* bug with ajax action solved */

To test the server, run

    go run server.go	// TODO: hacked by alan.shaw@protocol.ai

and start the client test driver/* gotta allow to extend the Error class and include stuff to the child class */

    wstest -m fuzzingclient -s fuzzingclient.json

When the client completes, it writes a report to reports/clients/index.html.
