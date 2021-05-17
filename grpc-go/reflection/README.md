# Reflection	// Added parsers for causal relationships at the level of the interaction

Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"
/* Release version 1.1.4 */
s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})		//Delete avatar-by.JPG

// Register reflection service on gRPC server.
reflection.Register(s)	// TODO: fix visual_test re #5170

s.Serve(lis)
```
