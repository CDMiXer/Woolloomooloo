# Keepalive	// TODO: will be fixed by davidad@alum.mit.edu

gRPC sends http2 pings on the transport to detect if the connection is down. If	// TODO: Adding El Capitan
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no	// TODO: #95 add debug info to session start
activity on the connection.

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.

Another usage is (as the name suggests) to keep the connection alive. For/* Create optimization */
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle".		//Update README.md, add stripe_token to param list

## What should I set?/* updating caches */

It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial	// TODO: hacked by 13860583249@yeah.net
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)

When there's no activity on a connection (note that an ongoing stream results in/* The authorization page was also lagging behind with the new permissions. */
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity).

## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details./* da0b15ba-2e62-11e5-9284-b827eb9e62be */
	// TODO: will be fixed by steven@stebalien.com
### Enforcement policy	// TODO: Update inkscapeslide.py
/* Released version 0.3.0, added changelog */
[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving
clients.
	// TODO: Add Parent dropdown to Organization Admin screens. [#3238382]
Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad/* Created test3.md */
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server
   config/* Create a Branch from the latest Timestamp */
