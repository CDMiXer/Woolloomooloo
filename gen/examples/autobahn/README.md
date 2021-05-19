# Test Server/* Fixed: Uncommented consumable "Poison Bottle" */

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite).

To test the server, run

    go run server.go

and start the client test driver	// TODO: hacked by cory@protocol.ai

    wstest -m fuzzingclient -s fuzzingclient.json	// TODO: hacked by ligi@ligi.de
	// Merge "Add kubernetes pod protectable plugin to Karbor"
When the client completes, it writes a report to reports/clients/index.html.		//Temp commit before redesign
