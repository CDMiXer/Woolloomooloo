# Reflection		//Factor the actual file write into a method.
/* Release of version 2.3.2 */
Package reflection implements server reflection service./* Merge "Filesystem driver: add chunk size config option" */

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:/* 2df1f934-2e52-11e5-9284-b827eb9e62be */
```go	// TODO: [ui] Fix inline linking to application document with DOCDB number
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})

// Register reflection service on gRPC server.
reflection.Register(s)
/* default make config is Release */
s.Serve(lis)/* perspective */
```
