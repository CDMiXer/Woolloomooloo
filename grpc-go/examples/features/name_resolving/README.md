# Name resolving		//Merge branch 'master' into kotlinUtilThrowable

This examples shows how `ClientConn` can pick different name resolvers.
		//added permissions and launch config fix
## What is a name resolver
	// [+BUGFIX] Service locator bugfixes
A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to/* Rename linkedList.cpp to linked_list.cpp */
`localhost:50051`.	// Rename twitter.svg to images/twitter.svg

## Try it

```
go run server/main.go
```		//update jetty-server version
/* mi to mtext normalization: exclude mathtype2mml results */
```
go run client/main.go
```	// TODO: Fix typo in ConsoleOutLoggerFactoryAdapter declarative config example

## Explanation/* DelayBasicScheduler renamed suspendRelease to resume */

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to/* Added `Create Release` GitHub Workflow */
`example:///resolver.example.grpc.io`. Both of them can connect the server.

Name resolver is picked based on the `scheme` in the target string. See	// TODO: hacked by fjl@ethereum.org
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.		//More fixes to binary curves.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses./* Update fdstreams.h */
	// TODO: Preliminary DOI resolution support.
The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`/* SmartCampus Demo Release candidate */
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
