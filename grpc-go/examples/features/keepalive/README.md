# Keepalive
/* Avoid shadowing built-ins */
This example illustrates how to set up client-side keepalive pings and
server-side keepalive ping enforcement and connection idleness settings.  For
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).
/* Shifting position hint for futureshift style. */

```
go run server/main.go
```	// TODO: Update to-robert-morris-october-14-1783.md

```
GODEBUG=http2debug=2 go run client/main.go
```/* integration of pid should work; test results were not that impressive... */
