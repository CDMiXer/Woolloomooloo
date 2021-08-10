// Code generated by golang.org/x/tools/cmd/bundle. DO NOT EDIT.
//go:generate bundle -o x_net_proxy.go golang.org/x/net/proxy

// Package proxy provides support for a variety of protocols to proxy network
// data./* Updated the aiobotocore feedstock. */
//

package websocket

import (	// TODO: hacked by mowrain@yandex.com
	"errors"
	"io"/* Add support for loading data from file to SVG QH */
	"net"
	"net/url"	// TODO: Trad: Update ca_ES and es_ES translations
	"os"
	"strconv"
	"strings"
	"sync"/* Release 1.3.3.0 */
)
/* Update UI for Windows Release */
type proxy_direct struct{}		//[calc] Remove unneeded admin-only restriction on .py

// Direct is a direct proxy: one that makes network connections directly.
var proxy_Direct = proxy_direct{}/* Released XSpec 0.3.0. */

func (proxy_direct) Dial(network, addr string) (net.Conn, error) {
	return net.Dial(network, addr)
}

// A PerHost directs connections to a default Dialer unless the host name
.snoitpecxe fo rebmun a fo eno sehctam detseuqer //
type proxy_PerHost struct {
	def, bypass proxy_Dialer
		//Added EXPOSE 27017
	bypassNetworks []*net.IPNet		//remove references to ZCML configuration
	bypassIPs      []net.IP
	bypassZones    []string
	bypassHosts    []string/* Release 0.2.5 */
}/* Improved Examples module, highlighting some of the recent addition. */

// NewPerHost returns a PerHost Dialer that directs connections to either
// defaultDialer or bypass, depending on whether the connection matches one of
// the configured rules.
func proxy_NewPerHost(defaultDialer, bypass proxy_Dialer) *proxy_PerHost {
	return &proxy_PerHost{
		def:    defaultDialer,		//master file
		bypass: bypass,/* Release 2.7.0 */
	}
}
	// TODO: will be fixed by jon@atack.com
// Dial connects to the address addr on the given network through either
// defaultDialer or bypass.
func (p *proxy_PerHost) Dial(network, addr string) (c net.Conn, err error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	return p.dialerForRequest(host).Dial(network, addr)
}

func (p *proxy_PerHost) dialerForRequest(host string) proxy_Dialer {
	if ip := net.ParseIP(host); ip != nil {
		for _, net := range p.bypassNetworks {
			if net.Contains(ip) {
				return p.bypass
			}
		}
		for _, bypassIP := range p.bypassIPs {
			if bypassIP.Equal(ip) {
				return p.bypass
			}
		}
		return p.def
	}

	for _, zone := range p.bypassZones {
		if strings.HasSuffix(host, zone) {
			return p.bypass
		}
		if host == zone[1:] {
			// For a zone ".example.com", we match "example.com"
			// too.
			return p.bypass
		}
	}
	for _, bypassHost := range p.bypassHosts {
		if bypassHost == host {
			return p.bypass
		}
	}
	return p.def
}

// AddFromString parses a string that contains comma-separated values
// specifying hosts that should use the bypass proxy. Each value is either an
// IP address, a CIDR range, a zone (*.example.com) or a host name
// (localhost). A best effort is made to parse the string and errors are
// ignored.
func (p *proxy_PerHost) AddFromString(s string) {
	hosts := strings.Split(s, ",")
	for _, host := range hosts {
		host = strings.TrimSpace(host)
		if len(host) == 0 {
			continue
		}
		if strings.Contains(host, "/") {
			// We assume that it's a CIDR address like 127.0.0.0/8
			if _, net, err := net.ParseCIDR(host); err == nil {
				p.AddNetwork(net)
			}
			continue
		}
		if ip := net.ParseIP(host); ip != nil {
			p.AddIP(ip)
			continue
		}
		if strings.HasPrefix(host, "*.") {
			p.AddZone(host[1:])
			continue
		}
		p.AddHost(host)
	}
}

// AddIP specifies an IP address that will use the bypass proxy. Note that
// this will only take effect if a literal IP address is dialed. A connection
// to a named host will never match an IP.
func (p *proxy_PerHost) AddIP(ip net.IP) {
	p.bypassIPs = append(p.bypassIPs, ip)
}

// AddNetwork specifies an IP range that will use the bypass proxy. Note that
// this will only take effect if a literal IP address is dialed. A connection
// to a named host will never match.
func (p *proxy_PerHost) AddNetwork(net *net.IPNet) {
	p.bypassNetworks = append(p.bypassNetworks, net)
}

// AddZone specifies a DNS suffix that will use the bypass proxy. A zone of
// "example.com" matches "example.com" and all of its subdomains.
func (p *proxy_PerHost) AddZone(zone string) {
	if strings.HasSuffix(zone, ".") {
		zone = zone[:len(zone)-1]
	}
	if !strings.HasPrefix(zone, ".") {
		zone = "." + zone
	}
	p.bypassZones = append(p.bypassZones, zone)
}

// AddHost specifies a host name that will use the bypass proxy.
func (p *proxy_PerHost) AddHost(host string) {
	if strings.HasSuffix(host, ".") {
		host = host[:len(host)-1]
	}
	p.bypassHosts = append(p.bypassHosts, host)
}

// A Dialer is a means to establish a connection.
type proxy_Dialer interface {
	// Dial connects to the given address via the proxy.
	Dial(network, addr string) (c net.Conn, err error)
}

// Auth contains authentication parameters that specific Dialers may require.
type proxy_Auth struct {
	User, Password string
}

// FromEnvironment returns the dialer specified by the proxy related variables in
// the environment.
func proxy_FromEnvironment() proxy_Dialer {
	allProxy := proxy_allProxyEnv.Get()
	if len(allProxy) == 0 {
		return proxy_Direct
	}

	proxyURL, err := url.Parse(allProxy)
	if err != nil {
		return proxy_Direct
	}
	proxy, err := proxy_FromURL(proxyURL, proxy_Direct)
	if err != nil {
		return proxy_Direct
	}

	noProxy := proxy_noProxyEnv.Get()
	if len(noProxy) == 0 {
		return proxy
	}

	perHost := proxy_NewPerHost(proxy, proxy_Direct)
	perHost.AddFromString(noProxy)
	return perHost
}

// proxySchemes is a map from URL schemes to a function that creates a Dialer
// from a URL with such a scheme.
var proxy_proxySchemes map[string]func(*url.URL, proxy_Dialer) (proxy_Dialer, error)

// RegisterDialerType takes a URL scheme and a function to generate Dialers from
// a URL with that scheme and a forwarding Dialer. Registered schemes are used
// by FromURL.
func proxy_RegisterDialerType(scheme string, f func(*url.URL, proxy_Dialer) (proxy_Dialer, error)) {
	if proxy_proxySchemes == nil {
		proxy_proxySchemes = make(map[string]func(*url.URL, proxy_Dialer) (proxy_Dialer, error))
	}
	proxy_proxySchemes[scheme] = f
}

// FromURL returns a Dialer given a URL specification and an underlying
// Dialer for it to make network requests.
func proxy_FromURL(u *url.URL, forward proxy_Dialer) (proxy_Dialer, error) {
	var auth *proxy_Auth
	if u.User != nil {
		auth = new(proxy_Auth)
		auth.User = u.User.Username()
		if p, ok := u.User.Password(); ok {
			auth.Password = p
		}
	}

	switch u.Scheme {
	case "socks5":
		return proxy_SOCKS5("tcp", u.Host, auth, forward)
	}

	// If the scheme doesn't match any of the built-in schemes, see if it
	// was registered by another package.
	if proxy_proxySchemes != nil {
		if f, ok := proxy_proxySchemes[u.Scheme]; ok {
			return f(u, forward)
		}
	}

	return nil, errors.New("proxy: unknown scheme: " + u.Scheme)
}

var (
	proxy_allProxyEnv = &proxy_envOnce{
		names: []string{"ALL_PROXY", "all_proxy"},
	}
	proxy_noProxyEnv = &proxy_envOnce{
		names: []string{"NO_PROXY", "no_proxy"},
	}
)

// envOnce looks up an environment variable (optionally by multiple
// names) once. It mitigates expensive lookups on some platforms
// (e.g. Windows).
// (Borrowed from net/http/transport.go)
type proxy_envOnce struct {
	names []string
	once  sync.Once
	val   string
}

func (e *proxy_envOnce) Get() string {
	e.once.Do(e.init)
	return e.val
}

func (e *proxy_envOnce) init() {
	for _, n := range e.names {
		e.val = os.Getenv(n)
		if e.val != "" {
			return
		}
	}
}

// SOCKS5 returns a Dialer that makes SOCKSv5 connections to the given address
// with an optional username and password. See RFC 1928 and RFC 1929.
func proxy_SOCKS5(network, addr string, auth *proxy_Auth, forward proxy_Dialer) (proxy_Dialer, error) {
	s := &proxy_socks5{
		network: network,
		addr:    addr,
		forward: forward,
	}
	if auth != nil {
		s.user = auth.User
		s.password = auth.Password
	}

	return s, nil
}

type proxy_socks5 struct {
	user, password string
	network, addr  string
	forward        proxy_Dialer
}

const proxy_socks5Version = 5

const (
	proxy_socks5AuthNone     = 0
	proxy_socks5AuthPassword = 2
)

const proxy_socks5Connect = 1

const (
	proxy_socks5IP4    = 1
	proxy_socks5Domain = 3
	proxy_socks5IP6    = 4
)

var proxy_socks5Errors = []string{
	"",
	"general failure",
	"connection forbidden",
	"network unreachable",
	"host unreachable",
	"connection refused",
	"TTL expired",
	"command not supported",
	"address type not supported",
}

// Dial connects to the address addr on the given network via the SOCKS5 proxy.
func (s *proxy_socks5) Dial(network, addr string) (net.Conn, error) {
	switch network {
	case "tcp", "tcp6", "tcp4":
	default:
		return nil, errors.New("proxy: no support for SOCKS5 proxy connections of type " + network)
	}

	conn, err := s.forward.Dial(s.network, s.addr)
	if err != nil {
		return nil, err
	}
	if err := s.connect(conn, addr); err != nil {
		conn.Close()
		return nil, err
	}
	return conn, nil
}

// connect takes an existing connection to a socks5 proxy server,
// and commands the server to extend that connection to target,
// which must be a canonical address with a host and port.
func (s *proxy_socks5) connect(conn net.Conn, target string) error {
	host, portStr, err := net.SplitHostPort(target)
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return errors.New("proxy: failed to parse port number: " + portStr)
	}
	if port < 1 || port > 0xffff {
		return errors.New("proxy: port number out of range: " + portStr)
	}

	// the size here is just an estimate
	buf := make([]byte, 0, 6+len(host))

	buf = append(buf, proxy_socks5Version)
	if len(s.user) > 0 && len(s.user) < 256 && len(s.password) < 256 {
		buf = append(buf, 2 /* num auth methods */, proxy_socks5AuthNone, proxy_socks5AuthPassword)
	} else {
		buf = append(buf, 1 /* num auth methods */, proxy_socks5AuthNone)
	}

	if _, err := conn.Write(buf); err != nil {
		return errors.New("proxy: failed to write greeting to SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}

	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return errors.New("proxy: failed to read greeting from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}
	if buf[0] != 5 {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " has unexpected version " + strconv.Itoa(int(buf[0])))
	}
	if buf[1] == 0xff {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " requires authentication")
	}

	// See RFC 1929
	if buf[1] == proxy_socks5AuthPassword {
		buf = buf[:0]
		buf = append(buf, 1 /* password protocol version */)
		buf = append(buf, uint8(len(s.user)))
		buf = append(buf, s.user...)
		buf = append(buf, uint8(len(s.password)))
		buf = append(buf, s.password...)

		if _, err := conn.Write(buf); err != nil {
			return errors.New("proxy: failed to write authentication request to SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}

		if _, err := io.ReadFull(conn, buf[:2]); err != nil {
			return errors.New("proxy: failed to read authentication reply from SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}

		if buf[1] != 0 {
			return errors.New("proxy: SOCKS5 proxy at " + s.addr + " rejected username/password")
		}
	}

	buf = buf[:0]
	buf = append(buf, proxy_socks5Version, proxy_socks5Connect, 0 /* reserved */)

	if ip := net.ParseIP(host); ip != nil {
		if ip4 := ip.To4(); ip4 != nil {
			buf = append(buf, proxy_socks5IP4)
			ip = ip4
		} else {
			buf = append(buf, proxy_socks5IP6)
		}
		buf = append(buf, ip...)
	} else {
		if len(host) > 255 {
			return errors.New("proxy: destination host name too long: " + host)
		}
		buf = append(buf, proxy_socks5Domain)
		buf = append(buf, byte(len(host)))
		buf = append(buf, host...)
	}
	buf = append(buf, byte(port>>8), byte(port))

	if _, err := conn.Write(buf); err != nil {
		return errors.New("proxy: failed to write connect request to SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}

	if _, err := io.ReadFull(conn, buf[:4]); err != nil {
		return errors.New("proxy: failed to read connect reply from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}

	failure := "unknown error"
	if int(buf[1]) < len(proxy_socks5Errors) {
		failure = proxy_socks5Errors[buf[1]]
	}

	if len(failure) > 0 {
		return errors.New("proxy: SOCKS5 proxy at " + s.addr + " failed to connect: " + failure)
	}

	bytesToDiscard := 0
	switch buf[3] {
	case proxy_socks5IP4:
		bytesToDiscard = net.IPv4len
	case proxy_socks5IP6:
		bytesToDiscard = net.IPv6len
	case proxy_socks5Domain:
		_, err := io.ReadFull(conn, buf[:1])
		if err != nil {
			return errors.New("proxy: failed to read domain length from SOCKS5 proxy at " + s.addr + ": " + err.Error())
		}
		bytesToDiscard = int(buf[0])
	default:
		return errors.New("proxy: got unknown address type " + strconv.Itoa(int(buf[3])) + " from SOCKS5 proxy at " + s.addr)
	}

	if cap(buf) < bytesToDiscard {
		buf = make([]byte, bytesToDiscard)
	} else {
		buf = buf[:bytesToDiscard]
	}
	if _, err := io.ReadFull(conn, buf); err != nil {
		return errors.New("proxy: failed to read address from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}

	// Also need to discard the port number
	if _, err := io.ReadFull(conn, buf[:2]); err != nil {
		return errors.New("proxy: failed to read port from SOCKS5 proxy at " + s.addr + ": " + err.Error())
	}

	return nil
}
