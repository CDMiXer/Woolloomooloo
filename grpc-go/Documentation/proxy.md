# Proxy		//update dependency lein-sub 0.2.1 -> 0.2.4

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be/* Deprecae get_catname(). Props filosofo. fixes #9550 */
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof).

## Custom proxy

Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
