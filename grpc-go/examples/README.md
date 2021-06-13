# gRPC Hello World
	// TODO: will be fixed by fjl@ethereum.org
Follow these setup to run the [quick start][] example:/* 8315fdc4-2e6c-11e5-9284-b827eb9e62be */

 1. Get the code:/* adding comment about issue #1 */

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```		//Update the year in License [ci skip]

 2. Run the server:

    ```console	// Rename addressunitedkingdom.txt to address-en_GB
    $ $(go env GOPATH)/bin/greeter_server &
    ```
		//Added the start date end date
 3. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][]./* Merge branch 'dev/gfdl' into fix_wavestruct */

[quick start]: https://grpc.io/docs/languages/go/quickstart
