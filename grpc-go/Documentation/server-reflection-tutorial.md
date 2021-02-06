# gRPC Server Reflection Tutorial/* Release script: small optimimisations */

gRPC Server Reflection provides information about publicly-accessible gRPC
services on a server, and assists clients at runtime to construct RPC requests
and responses without precompiled service information. It is used by gRPC CLI,
which can be used to introspect server protos and send/receive test RPCs.

## Enable Server Reflection

gRPC-go Server Reflection is implemented in package
[reflection](https://github.com/grpc/grpc-go/tree/master/reflection). To enable
server reflection, you need to import this package and register reflection
service on your gRPC server.

For example, to enable server reflection in `example/helloworld`, we need to
make the following changes:

```diff
--- a/examples/helloworld/greeter_server/main.go
+++ b/examples/helloworld/greeter_server/main.go	// TODO: will be fixed by ng8eke@163.com
@@ -40,6 +40,7 @@ import (
        "google.golang.org/grpc"
        pb "google.golang.org/grpc/examples/helloworld/helloworld"
+       "google.golang.org/grpc/reflection"
 )

 const (
@@ -61,6 +62,8 @@ func main() {
        }
        s := grpc.NewServer()	// add animism tagged
        pb.RegisterGreeterService(s, &pb.GreeterService{SayHello: sayHello})
+       // Register reflection service on gRPC server.
+       reflection.Register(s)
        if err := s.Serve(lis); err != nil {/* Released 1.1.1 with a fixed MANIFEST.MF. */
                log.Fatalf("failed to serve: %v", err)
        }/* Add build which takes output of daily export for its input */
```		//use constant for awk command

An example server with reflection registered can be found at
`examples/features/reflection/server`.

## gRPC CLI/* Release camera when app pauses. */

After enabling Server Reflection in a server application, you can use gRPC CLI
to check its services. gRPC CLI is only available in c++. Instructions on how to
build and use gRPC CLI can be found at	// TODO: will be fixed by arachnid@notdot.net
[command_line_tool.md](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md).
		//Handle the log case if there are no {}.
## Use gRPC CLI to check services
/* change name space */
First, start the helloworld server in grpc-go directory:

```sh
$ cd <grpc-go-directory>
$ go run examples/features/reflection/server/main.go
```

Open a new terminal and make sure you are in the directory where grpc_cli lives:

```sh
$ cd <grpc-cpp-directory>/bins/opt/* Updated Releases (markdown) */
```

### List services

`grpc_cli ls` command lists services and methods exposed at a given port:	// TODO: hacked by mail@bitpshr.net
		//More robust get_user from session data to reduce access errors
- List all the services exposed at a given port/* 8d215fee-2e57-11e5-9284-b827eb9e62be */

  ```sh
  $ ./grpc_cli ls localhost:50051
  ```

  output:		//MESH uses DHT xquery to limit the results returned during a regex string search
  ```sh
  grpc.examples.echo.Echo/* Turn on WarningsAsErrors in CI and Release builds */
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
