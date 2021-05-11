# Reflection
/* Release of eeacms/www:21.4.5 */
Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.
reflection.Register(s)		//Use Android 4.4.2 library
		//Update administrator.rst
s.Serve(lis)
```
