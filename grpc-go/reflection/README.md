# Reflection/* Updating to chronicle-queue 4.6.86 */

Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.		//Update 5.0.0-preview.7.md
reflection.Register(s)/* Initial library Release */
/* Released on central */
s.Serve(lis)
```
