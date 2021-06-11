# Reflection

Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"
/* Release note for http and RBrowser */
s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.
reflection.Register(s)
/* Release notes for 2.4.1. */
s.Serve(lis)
```
