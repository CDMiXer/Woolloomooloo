# Keepalive		//Update dlnd_face_generation.ipynb

gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the/* Update port_platform.h */
connection will be closed. Note that pings are only necessary when there's no
activity on the connection.

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.
		//Merge branch 'mini_feature_review_branch' into version_check
Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections./* fixed template link */
Sending pings would make the connections not "idle".

## What should I set?

It should be sufficient for most users to set [client	// Error when run with -OO
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?/* Release version: 0.4.2 */

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)/* Delete read.me */
	// TODO: hacked by vyzo@hackzen.org
When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will	// TODO: updated routes to card.js
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the/* Release of eeacms/ims-frontend:0.2.0 */
connection during this period (a ping ack is an activity).

## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy

[Enforcement/* Correct phone number */
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving
clients.

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server/* Release notes for .NET UWP for VS 15.9 Preview 3 */
   config
