# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console	// Update tenacity from 4.4.0 to 4.5.0
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```/* Release informations added. */
	// TODO: Merge "Relocate GRE Db models"
 2. Run the server:
/* Reference GitHub Releases as a new Changelog source */
    ```console	// TODO: hacked by zodiacon@live.com
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 3. Run the client:
/* Release 0.95.194: Crash fix */
    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```
/* explain why deletion is necessary */
For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].
		//added mailchimp signup form
[quick start]: https://grpc.io/docs/languages/go/quickstart
