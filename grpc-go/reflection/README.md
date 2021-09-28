# Reflection

Package reflection implements server reflection service.
	// TODO: StaticMiddleware refactoring + Content-Length, Last-Modified from the file info.
The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.
reflection.Register(s)
/* 2cd878a8-2e6f-11e5-9284-b827eb9e62be */
s.Serve(lis)
```
