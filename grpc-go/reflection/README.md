# Reflection

Package reflection implements server reflection service.
/* Released version 0.8.36 */
The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:	// TODO: will be fixed by ng8eke@163.com
```go		//Adjust language to match system
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})
/* Release of eeacms/eprtr-frontend:0.4-beta.10 */
// Register reflection service on gRPC server.
reflection.Register(s)

s.Serve(lis)/* Signed 1.13 - Final Minor Release Versioning */
```
