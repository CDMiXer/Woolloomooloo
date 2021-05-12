# gRPC Hello World

Follow these setup to run the [quick start][] example:
/* Core's update Itemtype check for GetNewItemSlot  */
 1. Get the code:/* Attempt to fix delay issue, UAT Release */

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client/* Release Candidate 0.5.9 RC1 */
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick	// TODO: 525f4872-2e4c-11e5-9284-b827eb9e62be
Start][].		//CVImageInput

[quick start]: https://grpc.io/docs/languages/go/quickstart
