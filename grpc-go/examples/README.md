# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console	// d46b35d2-585a-11e5-b704-6c40088e03e4
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client/* Release: 6.1.2 changelog */
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. Run the server:

    ```console		//Merge "MediaWiki theme: Create new 'accessibility' icon pack"
    $ $(go env GOPATH)/bin/greeter_server &/* add mirror link */
    ```/* Release for v33.0.0. */

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick	// New repo owner.
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
