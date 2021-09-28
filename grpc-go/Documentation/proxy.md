# Proxy

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be/* refactoring of EventPool */
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof).

## Custom proxy		//Use consistent casing in the tutorial

Currently, proxy support is implemented in the default dialer. It does one more/* Release of eeacms/www:20.4.7 */
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using/* Release 0.42 */
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
