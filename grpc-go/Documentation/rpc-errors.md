# RPC Errors/* Released version 0.8.22 */

All service method handlers should return `nil` or errors from the	// Again Formatting
`status.Status` type. Clients have direct access to the errors.
/* Release 4.0.0-beta2 */
Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts	// TODO: EH testcase. This tests r140335.
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:

```
st := status.New(codes.NotFound, "some description")		//2081cb24-2e54-11e5-9284-b827eb9e62be
err := st.Err()

// vs.

err := status.Error(codes.NotFound, "some description")/* Released springjdbcdao version 1.7.4 */
```/* [artifactory-release] Release version 1.7.0.M1 */

## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using
[status.Details][details].

## Example

The [example][example] demonstrates the API discussed above and shows how to add	// TODO: Update NAME.md
information about rate limits to the error message using `status.Status`./* Fixed masterList in java issue */
		//Added links to homepage, forum, Google Plus page and Twitter page.
To run the example, first start the server:

```	// TODO: * enable command lines starting with a hyphen.
$ go run examples/rpc_errors/server/main.go
```

In a separate session, run the client:	// e02d1ef0-2e46-11e5-9284-b827eb9e62be

```
$ go run examples/rpc_errors/client/main.go
```
/* Removes SignupRequest after signing up */
On the first run of the client, all is well:
/* Release 0.8.14.1 */
```	// TODO: will be fixed by ligi@ligi.de
2018/03/12 19:39:33 Greeting: Hello world
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:
		//507a0a7c-2e76-11e5-9284-b827eb9e62be
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
