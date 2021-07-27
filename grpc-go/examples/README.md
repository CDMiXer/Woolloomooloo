# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console/* a better set attribute for inkweb */
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```/* Fixes JS error when editing course template */

 2. Run the server:/* Release notes for 1.0.85 */

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client/* rev 727693 */
    Greeting: Hello world		//Rebuilt index with scortasa
    ```
/* Add lab1 readme */
For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick	// TODO: hacked by nagydani@epointsystem.org
Start][]./* Create mi_conexion_panel.php */

[quick start]: https://grpc.io/docs/languages/go/quickstart
