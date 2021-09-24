# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:/* Updated GUI to reduce nagging behaviour about editArea */
/* Syntax highlighting in the README for code blocks */
    ```console	// Moved Travis build image in README
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```/* Info for Release5 */

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```		//Add default to installed version

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart/* Added Travis Github Releases support to the travis configuration file. */
