# Keepalive

gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no		//a30f3c52-2e4f-11e5-9284-b827eb9e62be
activity on the connection./* Release version: 1.0.6 */

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?/* New Release of swak4Foam (with finiteArea) */

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner./* Pulled the counting functionality into the JsonElementCount object. */

Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections.
Sending pings would make the connections not "idle".

## What should I set?

It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams)./* Updated field names. Added new script.  */

## What will happen?

(The behavior described here is specific for gRPC-go, it might be slightly
different in other languages.)
/* Simplify JSON response step definition. */
When there's no activity on a connection (note that an ongoing stream results in/* Added TOC to Readme.md */
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity)./* Release 1.8.0 */

## What about server side?
	// Delete boxplot.R
Server has similar `Time` and `Timeout` settings as client. Server can also
configure connection max-age. See [server/* add Deus-Exiroze icon */
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.

### Enforcement policy

[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is
a special setting on server side to protect server from malicious or misbehaving
clients.

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings/* fix(package): update random-http-useragent to version 1.1.11 */
 - Client sends pings when there's no stream and this is disallowed by server
   config
