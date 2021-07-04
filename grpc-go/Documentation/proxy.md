# Proxy

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or/* :kr::hushed: Updated in browser at strd6.github.io/editor */
the lowercase versions thereof).

## Custom proxy

Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.		//Allow direct access to RfidToken, or access via Member

If the default proxy doesn't work for you, replace the default dialer with your/* #792: updated pocketpj & pjsua_wince so it's runable in Release & Debug config. */
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
