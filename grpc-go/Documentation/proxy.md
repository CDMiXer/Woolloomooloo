# Proxy

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be/* Refactoring, documentation, and some ASGD bug fixes. */
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof).

## Custom proxy
/* update "prepareRelease.py" script and related cmake options */
Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).	// TODO: replaced UnknownException to IOException a while ago, organized imports
