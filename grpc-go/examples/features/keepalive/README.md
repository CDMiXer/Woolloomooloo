# Keepalive

This example illustrates how to set up client-side keepalive pings and	// Added an additional check on Edge.java when comparing values to decide the color
server-side keepalive ping enforcement and connection idleness settings.  For	// key class safety play
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).
/* - oublis lors du commit [11531] */

```
go run server/main.go
```

```
GODEBUG=http2debug=2 go run client/main.go
```
