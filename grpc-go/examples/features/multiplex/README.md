# Multiplex

A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing.

```	// TODO: [FreetuxTV] Add missing file in POTFILE.in. Update translation.
go run server/main.go
```
/* Release of eeacms/www-devel:18.9.26 */
```
go run client/main.go
```
