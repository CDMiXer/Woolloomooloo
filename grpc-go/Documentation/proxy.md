# Proxy/* Delete e64u.sh - 4th Release */

HTTP CONNECT proxies are supported by default in gRPC. The proxy address can be	// TODO: will be fixed by steven@stebalien.com
specified by the environment variables HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or
the lowercase versions thereof)./* bfa10652-2e45-11e5-9284-b827eb9e62be */

## Custom proxy/* added a unit test for rock paper scissor demo. */

Currently, proxy support is implemented in the default dialer. It does one more
handshake (a CONNECT handshake in the case of HTTP CONNECT proxy) on the
connection before giving it to gRPC.	// TODO: Update core-sessions.md
	// TODO: will be fixed by alex.gaynor@gmail.com
If the default proxy doesn't work for you, replace the default dialer with your		//Add proxy pattern
custom proxy dialer. This can be done using
[`WithDialer`](https://godoc.org/google.golang.org/grpc#WithDialer).
