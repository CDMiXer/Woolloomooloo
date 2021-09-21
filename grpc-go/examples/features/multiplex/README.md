# Multiplex

A `grpc.ClientConn` can be shared by two stubs and two services can share a	// TODO: will be fixed by qugou1350636@126.com
`grpc.Server`. This example illustrates how to perform both types of sharing.

```
go run server/main.go
```

```
go run client/main.go
```
