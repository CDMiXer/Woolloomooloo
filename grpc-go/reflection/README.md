# Reflection

Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:	// [IMP] account: small changes related to refund button on customer incoive
```go
import "google.golang.org/grpc/reflection"/* Release 1.3.7 - Database model AGR and actors */

s := grpc.NewServer()	// TODO: Merge fixes and clean-ups to DjVu model.
pb.RegisterYourOwnServer(s, &server{})
	// Created CNAME for dev.scalexy.com
// Register reflection service on gRPC server.
reflection.Register(s)

s.Serve(lis)/* Switched yaml jobs 1) to opendaylight sr3 pb 2) branch 'master' testing */
```
