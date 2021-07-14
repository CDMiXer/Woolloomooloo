# gRPC Hello World

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. Run the server:
		//Typo in the read me
    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```
		//no need for this (nw)
 3. Run the client:

    ```console/* Beta Release (Tweaks and Help yet to be finalised) */
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```
/* Create tedt.c */
For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart		//Delete data_slides_BF_PH.js
