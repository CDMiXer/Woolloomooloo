# Keepalive
/* Preparing WIP-Release v0.1.39.1-alpha */
gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no
activity on the connection./* New language: Catalan. */

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular/* test_upload.py: hush pyflakes */
situation is when the TCP connection drops packets (including FIN). It would		//Exceptions thrown by PholksaurusLib are caught and ignored.
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.

Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle".

## What should I set?
/* [MOD] Testing: stress tests updated */
It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?	// Merge "vpxdec: Disable writing of Y4M header when --md5 argument is specified."

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)
/* Release version: 1.2.3 */
When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.	// TODO: specify environment on addFiles calls
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity).		//Alle menu items werkend

## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy
/* make right arrow the same */
[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving
clients.		//Update es6.json

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings/* Release 2.12.2 */
 - Client sends pings when there's no stream and this is disallowed by server
   config
