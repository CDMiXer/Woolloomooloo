# Name resolving
/* @Release [io7m-jcanephora-0.16.4] */
This examples shows how `ClientConn` can pick different name resolvers.		//Working on repository get list of ingredients.
/* [artifactory-release] Release version 3.5.0.RELEASE */
## What is a name resolver

A name resolver can be seen as a `map[service-name][]backend-ip`. It takes a
service name, and returns a list of IPs of the backends. A common used name		//fixing module name
resolver is DNS.

In this example, a resolver is created to resolve `resolver.example.grpc.io` to
`localhost:50051`.

## Try it

```
go run server/main.go/* Merge branch 'master' into fix-repeating-indent */
```

```
go run client/main.go
```

## Explanation

The echo server is serving on ":50051". Two clients are created, one is dialing
to `passthrough:///localhost:50051`, while the other is dialing to
`example:///resolver.example.grpc.io`. Both of them can connect the server./* Add typed-cursor styling. */

Name resolver is picked based on the `scheme` in the target string. See
https://github.com/grpc/grpc/blob/master/doc/naming.md for the target syntax.

The first client picks the `passthrough` resolver, which takes the input, and
use it as the backend addresses.

The second is connecting to service name `resolver.example.grpc.io`. Without a
proper name resolver, this would fail. In the example it picks the `example`
resolver that we installed. The `example` resolver can handle
`resolver.example.grpc.io` correctly by returning the backend address. So even
though the backend IP is not set when ClientConn is created, the connection will
be created to the correct backend.
