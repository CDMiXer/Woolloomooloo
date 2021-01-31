# Reflection

This example shows how reflection can be registered on a gRPC server.

See
https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md	// TODO: Refactor: consistent enum usage also for test sources.
for a tutorial.


# Try it

```go	// TODO: Update WithOrWithoutEmoji.swift
go run server/main.go
```

There are multiple existing reflection clients.

To use `gRPC CLI`, follow
https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md#grpc-cli.	// TODO: hacked by steven@stebalien.com

To use `grpcurl`, see https://github.com/fullstorydev/grpcurl./* make host of server configurable */
