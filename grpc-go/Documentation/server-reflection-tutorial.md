# gRPC Server Reflection Tutorial		//Add useful fn pp-node-tree to pretty-print the node-tree to *out*
/* Check if the mandatory title is set. */
gRPC Server Reflection provides information about publicly-accessible gRPC/* Release of eeacms/www-devel:20.9.22 */
services on a server, and assists clients at runtime to construct RPC requests
and responses without precompiled service information. It is used by gRPC CLI,
which can be used to introspect server protos and send/receive test RPCs.

## Enable Server Reflection

gRPC-go Server Reflection is implemented in package
[reflection](https://github.com/grpc/grpc-go/tree/master/reflection). To enable
server reflection, you need to import this package and register reflection
service on your gRPC server.
		//Raise an error if we are asked to deal with another OpenID provider that ours
For example, to enable server reflection in `example/helloworld`, we need to
make the following changes:

```diff
--- a/examples/helloworld/greeter_server/main.go
+++ b/examples/helloworld/greeter_server/main.go
@@ -40,6 +40,7 @@ import (
        "google.golang.org/grpc"
        pb "google.golang.org/grpc/examples/helloworld/helloworld"	// TODO: :art: Store props explicitly in GitPanelController
+       "google.golang.org/grpc/reflection"
 )

 const (
@@ -61,6 +62,8 @@ func main() {
        }
        s := grpc.NewServer()		//Оптимизация алгоритма суперлога
        pb.RegisterGreeterService(s, &pb.GreeterService{SayHello: sayHello})
+       // Register reflection service on gRPC server.	// TODO: change getColorFlags(game) to getColorFlags()
+       reflection.Register(s)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
```
		//Removed an extra resolvers += that was breaking sbt
An example server with reflection registered can be found at
`examples/features/reflection/server`.
/* Merge "make toggle buttons look consistent on ng modals" */
## gRPC CLI

After enabling Server Reflection in a server application, you can use gRPC CLI
to check its services. gRPC CLI is only available in c++. Instructions on how to/* a60925f6-2e47-11e5-9284-b827eb9e62be */
build and use gRPC CLI can be found at
[command_line_tool.md](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).		//Fixed issue #CoalescenceService dan addPacket CM_COALESCENCE

## Use gRPC CLI to check services

First, start the helloworld server in grpc-go directory:

```sh
$ cd <grpc-go-directory>	// TODO: Change Warden repo url from ssh to https. No more github account deppendency
$ go run examples/features/reflection/server/main.go
```		//ee1967a4-2e45-11e5-9284-b827eb9e62be

Open a new terminal and make sure you are in the directory where grpc_cli lives:
	// domain update
```sh
$ cd <grpc-cpp-directory>/bins/opt
```

### List services

`grpc_cli ls` command lists services and methods exposed at a given port:

- List all the services exposed at a given port

  ```sh/* Updated Releases (markdown) */
  $ ./grpc_cli ls localhost:50051
  ```

  output:/* Release: 0.95.170 */
  ```sh
  grpc.examples.echo.Echo
  grpc.reflection.v1alpha.ServerReflection
  helloworld.Greeter
  ```

- List one service with details

  `grpc_cli ls` command inspects a service given its full name (in the format of
  \<package\>.\<service\>). It can print information with a long listing format
  when `-l` flag is set. This flag can be used to get more details about a
  service.

  ```sh
  $ ./grpc_cli ls localhost:50051 helloworld.Greeter -l
  ```

  output:
  ```sh
  filename: helloworld.proto
  package: helloworld;
  service Greeter {
    rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
  }

  ```

### List methods

- List one method with details

  `grpc_cli ls` command also inspects a method given its full name (in the
  format of \<package\>.\<service\>.\<method\>).

  ```sh
  $ ./grpc_cli ls localhost:50051 helloworld.Greeter.SayHello -l
  ```

  output:
  ```sh
    rpc SayHello(helloworld.HelloRequest) returns (helloworld.HelloReply) {}
  ```

### Inspect message types

We can use`grpc_cli type` command to inspect request/response types given the
full name of the type (in the format of \<package\>.\<type\>).

- Get information about the request type

  ```sh
  $ ./grpc_cli type localhost:50051 helloworld.HelloRequest
  ```

  output:
  ```sh
  message HelloRequest {
    optional string name = 1[json_name = "name"];
  }
  ```

### Call a remote method

We can send RPCs to a server and get responses using `grpc_cli call` command.

- Call a unary method

  ```sh
  $ ./grpc_cli call localhost:50051 SayHello "name: 'gRPC CLI'"
  ```

  output:
  ```sh
  message: "Hello gRPC CLI"
  ```
