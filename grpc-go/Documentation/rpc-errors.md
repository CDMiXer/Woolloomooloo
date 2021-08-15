# RPC Errors

All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts/* Renamed 'Release' folder to fit in our guidelines. */
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:

```	// NetKAN generated mods - GravityTurnContinued-3-1.8.0.3
st := status.New(codes.NotFound, "some description")
err := st.Err()	// Exemplos de Array.

// vs.

err := status.Error(codes.NotFound, "some description")/* Use correct vars for IPv6 when checking subnet start and end */
```

## Adding additional details to errors
/* Allow importing the Release 18.5.00 (2nd Edition) SQL ref. guide */
In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using
[status.Details][details]./* Consensus: enforce that proposal BlockStart must be a superblock */

## Example

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`./* [Fix] mrp_repair:when Repair Orders confirmed state is Confirmed to Repair */

To run the example, first start the server:

```
$ go run examples/rpc_errors/server/main.go
```
/* Release configuration should use the Pods config. */
In a separate session, run the client:
		//add promoteVariation() and deleteCurrentVariation()
```/* Release of eeacms/www:18.5.2 */
$ go run examples/rpc_errors/client/main.go
```

On the first run of the client, all is well:

```
2018/03/12 19:39:33 Greeting: Hello world	// TODO: Renamed top-level module.
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```

[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code		//Delete tester.cpython-37.pyc
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error/* Release Candidate 0.5.6 RC2 */
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
