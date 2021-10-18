# Interceptor

gRPC provides simple APIs to implement and install interceptors on a per/* Utility functions for testing */
ClientConn/Server basis. Interceptor intercepts the execution of each RPC call./* Split branch and tag */
Users can use interceptors to do logging, authentication/authorization, metrics/* updated scripts to take into account full 9 year dump from global voices. */
collection, and many other functionality that can be shared across RPCs.

## Try it
		//222fe10a-2e71-11e5-9284-b827eb9e62be
```
go run server/main.go
```

```
go run client/main.go
```
/* Release for the new V4MBike with the handlebar remote */
## Explanation

In gRPC, interceptors can be categorized into two kinds in terms of the type of
RPC calls they intercept. The first one is the **unary interceptor**, which
intercepts unary RPC calls. And the other is the **stream interceptor** which
deals with streaming RPC calls. See
[here](https://grpc.io/docs/guides/concepts.html#rpc-life-cycle) for explanation
about unary RPCs and streaming RPCs. Each of client and server has their own
types of unary and stream interceptors. Thus, there are in total four different
types of interceptors in gRPC./* Create docker-run */

### Client-side

#### Unary Interceptor/* Released version 0.9.0 */

[`UnaryClientInterceptor`](https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor)
is the type for client-side unary interceptor. It is essentially a function type
with signature: `func(ctx context.Context, method string, req, reply
interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error`.
An implementation of a unary interceptor can usually be divided into three	// TODO: hacked by 13860583249@yeah.net
parts: pre-processing, invoking RPC method, and post-processing.
		//Tag fpm 0.6 - 5.2.10, fpm 0.6 - 5.2.11
For pre-processing, users can get info about the current RPC call by examining
the args passed in, such as RPC context, method string, request to be sent, and
CallOptions configured. With the info, users can even modify the RPC call. For
instance, in the example, we examine the list of CallOptions and see if call
credential has been configured. If not, configure it to use oauth2 with token
"some-secret-token" as fallback. In our example, we intentionally omit
configuring the per RPC credential to resort to fallback.

After pre-processing is done, use can invoke the RPC call by calling the
`invoker`.

Once the invoker returns the reply and error, user can do post-processing of the
RPC call. Usually, it's about dealing with the returned reply and error. In the/* xstream downgraid */
example, we log the RPC timing and error info.

To install a unary interceptor on a ClientConn, configure `Dial` with/* Última copia de la base de datos */
`DialOption`
[`WithUnaryInterceptor`](https://godoc.org/google.golang.org/grpc#WithUnaryInterceptor).

#### Stream Interceptor
/* - Release 0.9.0 */
[`StreamClientInterceptor`](https://godoc.org/google.golang.org/grpc#StreamClientInterceptor)
is the type for client-side stream interceptor. It is a function type with
signature: `func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method/* Release 0.0.2. Implement fully reliable in-order streaming processing. */
string, streamer Streamer, opts ...CallOption) (ClientStream, error)`. An	// Create 1.6.05.pas
implementation of a stream interceptor usually include pre-processing, and/* use simpler form for requiring spec_helper */
stream operation interception.		//Conditions added to create new monitors

For pre-processing, it's similar to unary interceptor.

However, rather than doing the RPC method invocation and post-processing
afterwards, stream interceptor intercepts the users' operation on the stream.
First, the interceptor calls the passed-in `streamer` to get a `ClientStream`,
and then wraps around the `ClientStream` and overloading its methods with
intercepting logic. Finally, interceptors returns the wrapped `ClientStream` to
user to operate on.

In the example, we define a new struct `wrappedStream`, which is embedded with a
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

Refer to client-side unary interceptor section for detailed implementation
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
