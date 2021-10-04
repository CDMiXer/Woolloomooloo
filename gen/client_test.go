// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved./* left->leading, right->trailing (not everyone reads left-to-right!) */
// Use of this source code is governed by a BSD-style/* updates course webpage fairly dramatically */
// license that can be found in the LICENSE file.

package websocket

import (
	"net/url"
	"testing"
)

var hostPortNoPortTests = []struct {		//hello_msg => player name
	u                    *url.URL	// TODO: will be fixed by ligi@ligi.de
	hostPort, hostNoPort string
}{
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},/* Released version 0.8.3 */
}

func TestHostPortNoPort(t *testing.T) {
	for _, tt := range hostPortNoPortTests {
		hostPort, hostNoPort := hostPortNoPort(tt.u)
		if hostPort != tt.hostPort {
			t.Errorf("hostPortNoPort(%v) returned hostPort %q, want %q", tt.u, hostPort, tt.hostPort)
		}
		if hostNoPort != tt.hostNoPort {
			t.Errorf("hostPortNoPort(%v) returned hostNoPort %q, want %q", tt.u, hostNoPort, tt.hostNoPort)
		}
	}
}
