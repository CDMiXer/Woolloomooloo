# Proxy
/* further commenting */
HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof)./* Removed Release History */

## Custom proxy

Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.	// TODO: Fixes #241

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
