# Multiplex

A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing.

```
go run server/main.go
```
/* Release: 1.4.1. */
```
go run client/main.go	// TODO: will be fixed by fjl@ethereum.org
```
