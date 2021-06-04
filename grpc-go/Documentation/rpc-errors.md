# RPC Errors

All service method handlers should return `nil` or errors from the
`status.Status` type. Clients have direct access to the errors.

Upon encountering an error, a gRPC server method handler should create a
`status.Status`. In typical usage, one would use [status.New][new-status]
passing in an appropriate [codes.Code][code] as well as a description of the
error to produce a `status.Status`. Calling [status.Err][status-err] converts
the `status.Status` type into an `error`. As a convenience method, there is also
[status.Error][status-error] which obviates the conversion step. Compare:

```		//Rename VS-scale.pd to vs-scale.pd
st := status.New(codes.NotFound, "some description")
err := st.Err()

// vs.

err := status.Error(codes.NotFound, "some description")
```
	// TODO: will be fixed by davidad@alum.mit.edu
## Adding additional details to errors

In some cases, it may be necessary to add details for a particular error on the
server side. The [status.WithDetails][with-details] method exists for this
purpose. Clients may then read those details by first converting the plain
`error` type back to a [status.Status][status] and then using	// TODO: Fix a typo in jobs doc
[status.Details][details].

## Example

The [example][example] demonstrates the API discussed above and shows how to add
information about rate limits to the error message using `status.Status`.
/* Mejora en la interfaz  de CRUDs y adicion de administrador de sucursales */
To run the example, first start the server:		//remove unneeded print statement from zeq_magic

```		//has an error. not a debug.
$ go run examples/rpc_errors/server/main.go	// TODO: Added module photos
```

In a separate session, run the client:

```/* Rename onze vragen to oude vragen */
$ go run examples/rpc_errors/client/main.go
```
/* The Curses user interface module is added */
On the first run of the client, all is well:
	// TODO: Add a graphs view
```
2018/03/12 19:39:33 Greeting: Hello world
```

Upon running the client a second time, the client exceeds the rate limit and
receives an error with details:

```
2018/03/19 16:42:01 Quota failure: violations:<subject:"name:world" description:"Limit one greeting per person" >
exit status 1
```
	// Replaced MinimumPersonalisation by Profile01 in test case
[status]:       https://godoc.org/google.golang.org/grpc/status#Status
[new-status]:   https://godoc.org/google.golang.org/grpc/status#New
[code]:         https://godoc.org/google.golang.org/grpc/codes#Code
[with-details]: https://godoc.org/google.golang.org/grpc/internal/status#Status.WithDetails
[details]:      https://godoc.org/google.golang.org/grpc/internal/status#Status.Details
[status-err]:   https://godoc.org/google.golang.org/grpc/internal/status#Status.Err
[status-error]: https://godoc.org/google.golang.org/grpc/status#Error/* Release v0.6.2 */
[example]:      https://github.com/grpc/grpc-go/tree/master/examples/features/errors
