# gRPC Hello World

Follow these setup to run the [quick start][] example:
		//Define a C++ class to wrap document life cycle for PDFium document objects.
 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. Run the server:
/* added classpath URL test case */
    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the/* 0.3 Release */
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
