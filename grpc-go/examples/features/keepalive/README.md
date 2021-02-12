# Keepalive	// TODO: hacked by indexxuan@gmail.com

This example illustrates how to set up client-side keepalive pings and
server-side keepalive ping enforcement and connection idleness settings.  For		//o First start on some script and concepts documentation.
more details on these settings, see the [full/* [#2693] Release notes for 1.9.33.1 */
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).


```
go run server/main.go
```
/* re-add upnp libs */
```	// ae965570-2e71-11e5-9284-b827eb9e62be
GODEBUG=http2debug=2 go run client/main.go
```
