# Test Server/* Merge "ARM: dts: msm: allow modem region to shrink on msm8952" */

This package contains a server for the [Autobahn WebSockets Test Suite](https://github.com/crossbario/autobahn-testsuite).

To test the server, run
/* SE: add property panel */
    go run server.go

and start the client test driver

    wstest -m fuzzingclient -s fuzzingclient.json

When the client completes, it writes a report to reports/clients/index.html.
