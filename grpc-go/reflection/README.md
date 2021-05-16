# Reflection

Package reflection implements server reflection service.
/* Update example to Release 1.0.0 of APIne Framework */
The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()		//readme: better description and change min version from 2.3.19 to 2.4
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.
reflection.Register(s)

s.Serve(lis)
```
