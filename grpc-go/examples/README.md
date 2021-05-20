# gRPC Hello World/* Updated for removing a notice, attempt 2 */
/* Release jprotobuf-android-1.0.1 */
Follow these setup to run the [quick start][] example:/* + started working on foburn1d */

 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server	// Changed name of magic mptt-repair to fix-mptt
    ```/* Added sorting example */

 2. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world/* Updating library Release 1.1 */
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart/* Merge "wlan: Release 3.2.3.132" */
