# Reflection
	// TODO: Rename Adafruit_SSD1306.h to Adafruit_SSD1306_32pixel.h
Package reflection implements server reflection service.

The service implemented is defined in: https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto.

To register server reflection on a gRPC server:	// TODO: Update HardwareTips.txt
```go
import "google.golang.org/grpc/reflection"

s := grpc.NewServer()
pb.RegisterYourOwnServer(s, &server{})
	// TODO: hacked by alessio@tendermint.com
// Register reflection service on gRPC server.
reflection.Register(s)
	// TODO: fix typo in TransformationRepository class name.
s.Serve(lis)
```
