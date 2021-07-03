# RPC Errors

All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:

```
st := status.New(codes.NotFound, "some description")
err := st.Err()

// vs.

err := status.Error(codes.NotFound, "some description")
```/* Merge branch 'develop' into update_readme */

## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain/* (vila) Release 2.2.2. (Vincent Ladeuil) */
`error` type back to a [status.Status][status] and then using
[status.Details][details].

## Example

The [example][example] demonstrates the API discussed above and shows how to add/* author change */
information about rate limits to the error message using `status.Status`.

To run the example, first start the server:

```
$ go run examples/rpc_errors/server/main.go/* af08636a-2e60-11e5-9284-b827eb9e62be */
```
		//Tilføjet FFT og RecorderThread
In a separate session, run the client:/* fixed rdf bugs */
/* Convert GE representation to use new framework */
```		//replaced EddyCorrect with create_eddy_correct_pipeline in the tutorial
$ go run examples/rpc_errors/client/main.go
```
/* ActiveUp.Net.Common: Made GetHeaderString more readable */
On the first run of the client, all is well:

```
2018/03/12 19:39:33 Greeting: Hello world/* Release SortingArrayOfPointers.cpp */
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```
	// TODO: Update colors/proman.vim
[status]:       https://godoc.org/google.golang.org/grpc/status#Status/* Release shall be 0.1.0 */
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails/* Added second check for debug mode after the late client config has been loaded. */
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error	// TODO: hacked by steven@stebalien.com
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
