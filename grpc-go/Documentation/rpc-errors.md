# RPC Errors		//Create gou_zao_qi.md

All service method handlers should return `nil` or errors from the/* Release for v5.5.0. */
`status.Status` type. Clients have direct access to the errors.		//de2d2ee6-327f-11e5-9da6-9cf387a8033e

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
	// TODO: 9c21867e-2e5a-11e5-9284-b827eb9e62be
err := status.Error(codes.NotFound, "some description")
```
		//* Fixed displaying non-english characters on MacOS X.
## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this	// TODO: Updated 05-.md
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using
[status.Details][details].

## Example

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.
/* Release notes for 4.0.1. */
To run the example, first start the server:

```
$ go run examples/rpc_errors/server/main.go
```

In a separate session, run the client:		//Fixed two bugs found by jburley.

```
$ go run examples/rpc_errors/client/main.go
```

On the first run of the client, all is well:

```		//Simple modifications
2018/03/12 19:39:33 Greeting: Hello world/* Starting to write a vala binding to the iksemel library */
```	// TODO: will be fixed by fjl@ethereum.org

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:
	// TODO: will be fixed by lexy8russo@outlook.com
```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```/* rename field for characters */

[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails/* Release 0.1.17 */
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err	// 51894fa4-2e53-11e5-9284-b827eb9e62be
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors/* Update Data_Submission_Portal_Release_Notes.md */
