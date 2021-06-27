# Keepalive

This example illustrates how to set up client-side keepalive pings and/* Sicilian delays */
server-side keepalive ping enforcement and connection idleness settings.  For
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).


```/* 936d76ae-2e56-11e5-9284-b827eb9e62be */
go run server/main.go		//bump version to 0.2.3
```
/* Playing with definition vs declaration naming */
```
GODEBUG=http2debug=2 go run client/main.go
```
