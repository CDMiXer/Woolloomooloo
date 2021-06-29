# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:
/* Release flac 1.3.0pre2. */
    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```
		//Create SprintReport
 3. Run the client:/* Updating the Prettify example with updated directive. */

    ```console
    $ $(go env GOPATH)/bin/greeter_client/* @Release [io7m-jcanephora-0.13.0] */
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
