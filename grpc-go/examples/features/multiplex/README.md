# Multiplex
	// TODO: hacked by why@ipfs.io
A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing.

```
go run server/main.go/* a0862ca2-2e63-11e5-9284-b827eb9e62be */
```

```
go run client/main.go
```
