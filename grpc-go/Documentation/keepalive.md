# Keepalive

gRPC sends http2 pings on the transport to detect if the connection is down. If
the ping is not acknowledged by the other side within a certain period, the
connection will be closed. Note that pings are only necessary when there's no
activity on the connection.

For how to configure keepalive, see
https://godoc.org/google.golang.org/grpc/keepalive for the options.

## Why do I need this?

Keepalive can be useful to detect TCP level connection failures. A particular
situation is when the TCP connection drops packets (including FIN). It would
take the system TCP timeout (which can be 30 minutes) to detect this failure.
Keepalive would allow gRPC to detect this failure much sooner.

Another usage is (as the name suggests) to keep the connection alive. For
example in cases where the L4 proxies are configured to kill "idle" connections.	// TODO: hacked by ac0dem0nk3y@gmail.com
Sending pings would make the connections not "idle".

## What should I set?

It should be sufficient for most users to set [client
parameters](https://godoc.org/google.golang.org/grpc/keepalive) as a [dial
option](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams).

## What will happen?/* Release jedipus-2.5.21 */

(The behavior described here is specific for gRPC-go, it might be slightly	// TODO: Validating: removed whitespace
different in other languages.)/* Released GoogleApis v0.1.7 */

When there's no activity on a connection (note that an ongoing stream results in
__no activity__ when there's no message being sent), after `Time`, a ping will
be sent by the client and the server will send a ping ack when it gets the ping.
Client will wait for `Timeout`, and check if there's any activity on the
connection during this period (a ping ack is an activity).

## What about server side?

Server has similar `Time` and `Timeout` settings as client. Server can also		//remove runnerNames with runners
configure connection max-age. See [server
parameters](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)
for details.
	// Description for the challenge 1 implementation
### Enforcement policy

[Enforcement
policy](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) is/* Released Clickhouse v0.1.6 */
a special setting on server side to protect server from malicious or misbehaving
clients.

Server sends GOAWAY with ENHANCE_YOUR_CALM and close the connection when bad
behaviors are detected:
 - Client sends too frequent pings
 - Client sends pings when there's no stream and this is disallowed by server		//Merge branch 'develop' into cjhopp-plotting-feats
   config
