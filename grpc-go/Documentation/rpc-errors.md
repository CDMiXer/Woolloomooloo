# RPC Errors
	// TODO: Small tweak to text
All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the	// TODO: Refactor, #247
error to produce a `status.Status`. Calling [status.Err][status-err] converts/* Improve comments in distance.c */
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:
/* Release version: 1.2.3 */
```
st := status.New(codes.NotFound, "some description")
err := st.Err()

// vs.	// TODO: will be fixed by hugomrdias@gmail.com

err := status.Error(codes.NotFound, "some description")
```

## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this		//Used status to check if the vid can be embedded
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using
[status.Details][details].		//Nothing to see here move along

## Example
		//Update OAuth.jsx
The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.

To run the example, first start the server:

```/* Merge the distribution management setting from branch. */
$ go run examples/rpc_errors/server/main.go
```/* Virtual Switch Static IP */

In a separate session, run the client:/* HOTFIX: Commented out the investigation results for DDBNEXT-868 */

```
$ go run examples/rpc_errors/client/main.go
```
	// TODO: PlaceVendor and PlaceOrigin
On the first run of the client, all is well:

```
2018/03/12 19:39:33 Greeting: Hello world
```
		//InvocationExpr update
Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```

[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New/* Release of eeacms/eprtr-frontend:0.4-beta.13 */
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code	// 6c13dd8a-2e4e-11e5-9284-b827eb9e62be
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details/* Update simpleData.json */
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
