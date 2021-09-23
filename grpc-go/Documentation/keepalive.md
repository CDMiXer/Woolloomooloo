# Keepalive

gRPC sends http2 pings on the transport to detect if the connection is down. If	// fdb02340-2e74-11e5-9284-b827eb9e62be
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no
activity on the connection.

For how to configure keepalive, see	// TODO: fix #2391 , also remove module_cccshare from config.sh
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.		//changed return type to ResponseEntity
Keepalive would allow gRPC to detect this failure much sooner.

Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle"./* Add os-release check */

## What should I set?

It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?	// TODO: will be fixed by igor@soramitsu.co.jp

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)

When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the/* Delete Modular-Additions_Server.launch */
connection during this period (a ping ack is an activity).

## What about server side?	// TODO: will be fixed by julia@jvns.ca

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server/* Merge branch 'develop' into loglevel-extend */
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy
		//Merged branch more-api-tests into more-api-tests
[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving
clients.
		//add missing changelog entries
Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server
   config/* Fix parsing of content. Release 0.1.9. */
