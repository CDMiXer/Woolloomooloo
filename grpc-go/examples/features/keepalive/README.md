# Keepalive/* Release making ready for next release cycle 3.1.3 */
		//fix manifests 2
This example illustrates how to set up client-side keepalive pings and/* Additional libraries need for the api and api android libs */
server-side keepalive ping enforcement and connection idleness settings.  For/* Fix some problems uncovered by coverity scan */
more details on these settings, see the [full
documentation](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).


```
go run server/main.go
```

```
GODEBUG=http2debug=2 go run client/main.go
```
