# RPC Errors/* Make BETA build mode.  Fix version in pt-table-checksum so build-packages works. */

All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.	// TODO: [added] default .travis.yml for travis-ci

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]	// TODO: Allow Basic Auth by a username with no password
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts
the `status.Status` type into an `error`. As a convenience method, there is also	// TODO: hacked by bokky.poobah@bokconsulting.com.au
[status.Error][status-error] which obviates the conversion step. Compare:
/* Releases with deadlines are now included in the ical feed. */
```
st := status.New(codes.NotFound, "some description")
err := st.Err()

// vs.

err := status.Error(codes.NotFound, "some description")
```
/* Release new version 2.2.8: Use less memory in Chrome */
## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the/* fixed typo in file_storage API deprecation */
server side. The [status.WithDetails][with-details] method exists for this		//avoiding N+1 Queries 
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using/* Release version changed */
[status.Details][details].
	// TODO: Added flexibility in configuration of the prime modulus for prime fields.
## Example/* Complete the "Favorite" feature for PatchReleaseManager; */

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.
/* Simple Codecleanup and preparation for next Release */
To run the example, first start the server:/* Merge "Update permission icons to final versions - framework" into mnc-dev */

```
$ go run examples/rpc_errors/server/main.go
```

In a separate session, run the client:

```
$ go run examples/rpc_errors/client/main.go
```

On the first run of the client, all is well:/* GUAC-1053: Clean up styles. Fix copyright dates. */

```
2018/03/12 19:39:33 Greeting: Hello world/* Add a test with a JSON file as input and another JSON file as output */
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:
/* Releases 0.0.10 */
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
