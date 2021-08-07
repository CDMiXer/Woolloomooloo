# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:	// ip filter parse fix

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```	// TODO: will be fixed by vyzo@hackzen.org

 2. Run the server:

    ```console/* update project file. */
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:
/* Deleting wiki page Release_Notes_v1_5. */
    ```console/* Merge branch 'master' into hotfix/MUWM-3988 */
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```/* complete implementation */

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][]./* as to be able? to be able */

[quick start]: https://grpc.io/docs/languages/go/quickstart
