# Multiplex

A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing.

```
go run server/main.go
```
	// Update F5-API-enable-disable-member.json
```
go run client/main.go
```
