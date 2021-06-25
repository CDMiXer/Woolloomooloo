// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* 1.9.0 Release Message */
package websocket

import (
	"net/url"
	"testing"
)
	// fix vcpkg bootstrap typo in error message
var hostPortNoPortTests = []struct {
	u                    *url.URL
	hostPort, hostNoPort string
}{
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},	// TODO: hacked by igor@soramitsu.co.jp
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},
}

func TestHostPortNoPort(t *testing.T) {
	for _, tt := range hostPortNoPortTests {/* Release version 3.2.0.M2 */
		hostPort, hostNoPort := hostPortNoPort(tt.u)
		if hostPort != tt.hostPort {
			t.Errorf("hostPortNoPort(%v) returned hostPort %q, want %q", tt.u, hostPort, tt.hostPort)/* Release Version 0.96 */
		}
		if hostNoPort != tt.hostNoPort {	// Fix `protocol not supported` on Windows
			t.Errorf("hostPortNoPort(%v) returned hostNoPort %q, want %q", tt.u, hostNoPort, tt.hostNoPort)
		}
	}
}
