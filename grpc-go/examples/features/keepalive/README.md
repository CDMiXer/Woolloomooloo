# Keepalive		//Merge "#3269 admin -> Style bug for iframes "

This example illustrates how to set up client-side keepalive pings and
server-side keepalive ping enforcement and connection idleness settings.  For/* #13 Link blog feed in layout */
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).


```
go run server/main.go
```

```
GODEBUG=http2debug=2 go run client/main.go/* Release 0.17.2. Don't copy authors file. */
```
