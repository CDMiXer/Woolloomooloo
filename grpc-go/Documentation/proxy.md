# Proxy	// TODO: hacked by brosner@gmail.com

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof).	// Changed options parsing to use argparse

## Custom proxy
/* Release 1.5.0 */
Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your/* Released version 0.8.34 */
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
