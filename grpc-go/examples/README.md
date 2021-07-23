# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:		//Removed TODOs from README.md as they are complete

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client	// TODO: will be fixed by ac0dem0nk3y@gmail.com
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```
	// TODO: hacked by xaber.twt@gmail.com
 2. Run the server:	// TODO: will be fixed by alex.gaynor@gmail.com
/* Release of iText 5.5.11 */
    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```	// TODO: readme and examples

 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```
	// Add Movie Support to Speed.cd
For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick		//Create dataset-3.md
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
