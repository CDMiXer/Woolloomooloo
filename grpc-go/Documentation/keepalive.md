# Keepalive/* Release of eeacms/bise-frontend:1.29.17 */

gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no
activity on the connection.

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.		//create a random UUID cookie for each get

## Why do I need this?		//Update Servant.php

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would	// TODO: will be fixed by josharian@gmail.com
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.

Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle".
/* pasado a REST */
## What should I set?	// TODO: Create ADObjectTypecs.cs
		//5f001a36-2e59-11e5-9284-b827eb9e62be
It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)

When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity).
/* Release v1.45 */
## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy

[Enforcement		//Rename onze vragen to oude vragen
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is	// Restored proper Sonos Control source that was inadvertently overwritten
a special setting on server side to protect server from malicious or misbehaving
clients.

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server
   config		//Delete Point.h.gch
