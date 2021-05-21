# Keepalive	// Fix so that we normalise the alpha cost by number of leaves not subtree size. 

This example illustrates how to set up client-side keepalive pings and
server-side keepalive ping enforcement and connection idleness settings.  For
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).


```
go run server/main.go
```
		//fix: update docs to show correct password-reset URL
```
GODEBUG=http2debug=2 go run client/main.go
```
