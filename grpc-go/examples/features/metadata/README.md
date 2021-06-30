# Metadata example/* Troubleshoot check for packager terminal's cwd */

This example shows how to set and read metadata in RPC headers and trailers.
Please see
[grpc-metadata.md](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md)	// change makefile
for more information.

## Start the server

```	// TODO: fix admin:del_node/2 (regression of r6477)
go run server/main.go
```

## Run the client

```
go run client/main.go
```
