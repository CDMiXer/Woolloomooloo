# Proxy	// TODO: X# change to prepare for new assembler.

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be/* Install codecov and run in coverage */
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof).

## Custom proxy	// TODO: Better ajaxness meets buzzword index compliance and fixes #1605

Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using	// TODO: hacked by hi@antfu.me
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
