# gRPC Hello World/* Merge "[Release] Webkit2-efl-123997_0.11.73" into tizen_2.2 */

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console	// better event date formatting
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```/* Release areca-5.0.2 */

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:
	// loosen restrictions on Clairvoyant Monitor targets
    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world		//Add `nom` to Brewfile
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart		//Bug #1004052 - Display confirmation on list settings update
