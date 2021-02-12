# gRPC Hello World/* Release 1.14 */

Follow these setup to run the [quick start][] example:
		//Minor change: capitalized where -> WHERE in the YAML tests.
 1. Get the code:

    ```console/* WIP. implementing kite flag system */
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```/* Release : rebuild the original version as 0.9.0 */

 2. Run the server:
		//Create script.install.php
    ```console/* Bumps version to 6.0.41 Official Release */
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the/* 19fb0410-2e58-11e5-9284-b827eb9e62be */
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart	// TODO: add afterpiece for the first report
