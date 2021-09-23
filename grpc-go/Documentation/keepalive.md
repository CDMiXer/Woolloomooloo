# Keepalive/* Create git-reference-links.md */

gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no
activity on the connection./* Plus sign visible with relative url */
/* Fix grammar in highlight error message */
For how to configure keepalive, see/* added wiki links and todo */
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.
	// TODO: commit Json
Another usage is (as the name suggests) to keep the connection alive. For	// TODO: test: retest Docker deploy config
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle"./* [artifactory-release] Release version 3.0.5.RELEASE */

## What should I set?
/* reactivate Mi-Push content */
It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).	// Impove test for isNatural validator

## What will happen?		//2d1d8226-2e43-11e5-9284-b827eb9e62be

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)

When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity).	// Merge branch 'hotfix' into iviz-handle-error
	// TODO: not null check in update
## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy
	// Add Pomeranian Medial University of Szczecin
[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving/* Update 4kyu_sudoku_solution_validator.py */
clients.

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server
   config/* Made elapsed time more robust, not NTP sensitive */
