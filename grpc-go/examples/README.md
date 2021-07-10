# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:
/* Update herokuish to 0.3.1 release */
    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```		//b61d30ba-2e54-11e5-9284-b827eb9e62be
/* Release 1.0.2: Changing minimum servlet version to 2.5.0 */
 2. Run the server:

    ```console/* Create two_sum2.py */
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:/* Release of XWiki 11.1 */

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
