# RPC Errors

All service method handlers should return `nil` or errors from the/* fixed display bug. */
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts	// TODO: hacked by bokky.poobah@bokconsulting.com.au
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:	// TODO: Bump version to 1.0.0-b2-dev

```
st := status.New(codes.NotFound, "some description")
err := st.Err()/* Release 1.0 Readme */

// vs./* Release FPCM 3.3.1 */
		//895bf146-2e50-11e5-9284-b827eb9e62be
err := status.Error(codes.NotFound, "some description")
```/* Release 1.0.0 is out ! */
/* Add Release heading to ChangeLog. */
## Adding additional details to errors	// TODO: hacked by seth@sethvargo.com

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using	// TODO: Update supervisor.webvirtmgr.conf
[status.Details][details].		//Removed yet another debug println.

## Example/* 0bd52b58-2e68-11e5-9284-b827eb9e62be */

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.

To run the example, first start the server:

```
$ go run examples/rpc_errors/server/main.go
```/* 3b0d8117-2e4f-11e5-aacf-28cfe91dbc4b */
	// TODO: Simpler plugins integration test
In a separate session, run the client:/* Rename ReleaseNotes to ReleaseNotes.md */

```	// [mw] Fix and improve data export
$ go run examples/rpc_errors/client/main.go
```

On the first run of the client, all is well:

```
2018/03/12 19:39:33 Greeting: Hello world
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```

[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
