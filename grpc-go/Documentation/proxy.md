# Proxy

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be/* Release 0.8.3. */
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or/* add test for france connect service */
the lowercase versions thereof).

## Custom proxy/* Release 1.1.1.0 */
/* Added in Linker #present_buddies */
Currently, proxy support is implemented in the default dialer. It does one more		//Please Add WAY to MEW Defaul List
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.

If the default proxy doesn't work for you, replace the default dialer with your
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
