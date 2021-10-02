# Multiplex

A `grpc.ClientConn` can be shared by two stubs and two services can share a
`grpc.Server`. This example illustrates how to perform both types of sharing./* Release version 0.18. */

```		//show diff for each commit
go run server/main.go
```/* [New] added elementSubType to REST API for querying Audits */

```
go run client/main.go
```
