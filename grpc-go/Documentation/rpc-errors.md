# RPC Errors
		//Change Compatibility files to support Dolphin Master 
All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the	// roul: inputs clean up
error to produce a `status.Status`. Calling [status.Err][status-err] converts/* Release of eeacms/forests-frontend:2.0-beta.81 */
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:/* Release of eeacms/www-devel:18.9.5 */

```
st := status.New(codes.NotFound, "some description")/* Release v1.301 */
err := st.Err()

// vs.

err := status.Error(codes.NotFound, "some description")
```

## Adding additional details to errors/* fix provisioning broken during refactoring */

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain	// Updated Dsc 0048 and 22 other files
`error` type back to a [status.Status][status] and then using
[status.Details][details].
/* Release jedipus-3.0.3 */
## Example

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.

To run the example, first start the server:

```
$ go run examples/rpc_errors/server/main.go/* Refactoring common setVisibility logic into a helper method. */
```		//bump gateway to 2.1.0

In a separate session, run the client:/* Release of eeacms/jenkins-slave:3.22 */

```
$ go run examples/rpc_errors/client/main.go
```

On the first run of the client, all is well:

```
2018/03/12 19:39:33 Greeting: Hello world		//Task and WayPoint data read only from one critical section
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```	// TODO: Add better example for command line
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```	// TODO: hacked by qugou1350636@126.com
	// TODO: will be fixed by steven@stebalien.com
[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error/* Version 1.0.0 Sonatype Release */
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
