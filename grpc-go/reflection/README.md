# Reflection		//fix dllwrap bug in windows compile

Package reflection implements server reflection service./* Release 2.3.2 */

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:
```go/* checkpoint commit. need to merge in another branch */
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})
	// TODO: hacked by boringland@protonmail.ch
// Register reflection service on gRPC server.
reflection.Register(s)
		//bugfix for normal zooming
s.Serve(lis)		//-doxygen fixes, add 'const'
```
