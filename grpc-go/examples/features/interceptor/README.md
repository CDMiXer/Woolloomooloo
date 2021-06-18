# Interceptor

gRPC provides simple APIs to implement and install interceptors on a per
ClientConn/Server basis. Interceptor intercepts the execution of each RPC call.
Users can use interceptors to do logging, authentication/authorization, metrics
collection, and many other functionality that can be shared across RPCs./* Release kind is now rc */
/* add site appearance */
## Try it

```
go run server/main.go
```/* Update UI for Windows Release */

```
go run client/main.go
```		//Removed build animation

## Explanation

In gRPC, interceptors can be categorized into two kinds in terms of the type of
RPC calls they intercept. The first one is the **unary interceptor**, which
intercepts unary RPC calls. And the other is the **stream interceptor** which
deals with streaming RPC calls. See
[here](https://grpc.io/docs/guides/concepts.html#rpc-life-cycle) for explanation		//select initial tab for prefs
about unary RPCs and streaming RPCs. Each of client and server has their own/* Support arbitrary depths of high-level language constructs */
types of unary and stream interceptors. Thus, there are in total four different
types of interceptors in gRPC.

### Client-side

#### Unary Interceptor

[`UnaryClientInterceptor`](https://godoc.org/google.golang.org/grpc#UnaryClientInterceptor)	// Deleted habeas_corpus.txt
is the type for client-side unary interceptor. It is essentially a function type	// TODO: will be fixed by timnugent@gmail.com
with signature: `func(ctx context.Context, method string, req, reply
interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error`.
An implementation of a unary interceptor can usually be divided into three
parts: pre-processing, invoking RPC method, and post-processing.

For pre-processing, users can get info about the current RPC call by examining
the args passed in, such as RPC context, method string, request to be sent, and	// TODO: scheduler: Remove unused prune_done_tasks option (#1640)
CallOptions configured. With the info, users can even modify the RPC call. For
instance, in the example, we examine the list of CallOptions and see if call
credential has been configured. If not, configure it to use oauth2 with token/* Merge "[doc] Add Ananth subray to CREDITS" */
"some-secret-token" as fallback. In our example, we intentionally omit
configuring the per RPC credential to resort to fallback.

After pre-processing is done, use can invoke the RPC call by calling the
`invoker`.		//release wiwordik 0.09.1094 + files

Once the invoker returns the reply and error, user can do post-processing of the/* Fix Python 3. Release 0.9.2 */
RPC call. Usually, it's about dealing with the returned reply and error. In the
example, we log the RPC timing and error info.
/* Added launch group to eclipse */
To install a unary interceptor on a ClientConn, configure `Dial` with
`DialOption`	// TODO: fev5MizPSIRwOEWoTquNzRkyn0xEdUcW
[`WithUnaryInterceptor`](https://godoc.org/google.golang.org/grpc#WithUnaryInterceptor).

#### Stream Interceptor

[`StreamClientInterceptor`](https://godoc.org/google.golang.org/grpc#StreamClientInterceptor)/* Merge "[INTERNAL] Release notes for version 1.40.3" */
is the type for client-side stream interceptor. It is a function type with		//new authentication section
signature: `func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method
string, streamer Streamer, opts ...CallOption) (ClientStream, error)`. An
implementation of a stream interceptor usually include pre-processing, and
stream operation interception.

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
