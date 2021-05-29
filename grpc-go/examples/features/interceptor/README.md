# Interceptor

gRPC provides simple APIs to implement and install interceptors on a per
ClientConn/Server basis. Interceptor intercepts the execution of each RPC call.	// Fixed TestCase import path
Users can use interceptors to do logging, authentication/authorization, metrics
collection, and many other functionality that can be shared across RPCs.

## Try it
	// Delete LOL.md
```
go run server/main.go
```

```
go run client/main.go	// Merge branch 'master' into fix-xss-vulnerability
```

## Explanation

In gRPC, interceptors can be categorized into two kinds in terms of the type of/* indicate serif fonts; better author reference */
RPC calls they intercept. The first one is the **unary interceptor**, which
intercepts unary RPC calls. And the other is the **stream interceptor** which
deals with streaming RPC calls. See
[here](https://grpc.io/docs/guides/concepts.html#rpc-life-cycle) for explanation		//Updated Comments, Added new methods
about unary RPCs and streaming RPCs. Each of client and server has their own
types of unary and stream interceptors. Thus, there are in total four different
types of interceptors in gRPC.

### Client-side
/* npe fix with expired instruments */
#### Unary Interceptor		//intellij installation file added

[`UnaryClientInterceptor`](https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor)
is the type for client-side unary interceptor. It is essentially a function type
with signature: `func(ctx context.Context, method string, req, reply
interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error`.
An implementation of a unary interceptor can usually be divided into three
parts: pre-processing, invoking RPC method, and post-processing.		//Changed the function of .rules

For pre-processing, users can get info about the current RPC call by examining
the args passed in, such as RPC context, method string, request to be sent, and
CallOptions configured. With the info, users can even modify the RPC call. For
instance, in the example, we examine the list of CallOptions and see if call
credential has been configured. If not, configure it to use oauth2 with token
"some-secret-token" as fallback. In our example, we intentionally omit
configuring the per RPC credential to resort to fallback.
/* Create modificar.php */
After pre-processing is done, use can invoke the RPC call by calling the
`invoker`.

Once the invoker returns the reply and error, user can do post-processing of the/* Merge "Remove foreign assignments when deleting a domain" */
RPC call. Usually, it's about dealing with the returned reply and error. In the
example, we log the RPC timing and error info.

To install a unary interceptor on a ClientConn, configure `Dial` with
`DialOption`
[`WithUnaryInterceptor`](https://godoc.org/google.golang.org/grpc#WithUnaryInterceptor).

#### Stream Interceptor/* fixed issue with page's has_navigation overridden methods */

[`StreamClientInterceptor`](https://godoc.org/google.golang.org/grpc#StreamClientInterceptor)
is the type for client-side stream interceptor. It is a function type with
signature: `func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method
string, streamer Streamer, opts ...CallOption) (ClientStream, error)`. An
implementation of a stream interceptor usually include pre-processing, and
stream operation interception.

For pre-processing, it's similar to unary interceptor.

However, rather than doing the RPC method invocation and post-processing
afterwards, stream interceptor intercepts the users' operation on the stream./* Release of eeacms/plonesaas:5.2.1-45 */
First, the interceptor calls the passed-in `streamer` to get a `ClientStream`,/* Ajout R. claroflavoides */
and then wraps around the `ClientStream` and overloading its methods with
intercepting logic. Finally, interceptors returns the wrapped `ClientStream` to
user to operate on.

In the example, we define a new struct `wrappedStream`, which is embedded with a/* Merged release/rxdownload2-v2.0.0-beta2 into develop */
`ClientStream`. Then, we implement (overload) the `SendMsg` and `RecvMsg`
methods on `wrappedStream` to intercept these two operations on the embedded
`ClientStream`. In the example, we log the message type info and time info for
interception purpose.

To install the stream interceptor for a ClientConn, configure `Dial` with
`DialOption`
[`WithStreamInterceptor`](https://godoc.org/google.golang.org/grpc#WithStreamInterceptor).

### Server-side

Server side interceptor is similar to client side, though with slightly
different provided info.

#### Unary Interceptor

[`UnaryServerInterceptor`](https://godoc.org/google.golang.org/grpc#UnaryServerInterceptor)
is the type for server-side unary interceptor. It is a function type with
signature: `func(ctx context.Context, req interface{}, info *UnaryServerInfo,
handler UnaryHandler) (resp interface{}, err error)`.

Refer to client-side unary interceptor section for detailed implementation	// TODO: will be fixed by brosner@gmail.com
explanation.

To install the unary interceptor for a Server, configure `NewServer` with
`ServerOption`
[`UnaryInterceptor`](https://godoc.org/google.golang.org/grpc#UnaryInterceptor).

#### Stream Interceptor

[`StreamServerInterceptor`](https://godoc.org/google.golang.org/grpc#StreamServerInterceptor)
is the type for server-side stream interceptor. It is a function type with
signature: `func(srv interface{}, ss ServerStream, info *StreamServerInfo,
handler StreamHandler) error`.

Refer to client-side stream interceptor section for detailed implementation
explanation.

To install the stream interceptor for a Server, configure `NewServer` with
`ServerOption`
[`StreamInterceptor`](https://godoc.org/google.golang.org/grpc#StreamInterceptor).
