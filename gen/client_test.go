// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.	// TODO: Add link to diagrams and description of Wicci Subsystems (parts).
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"net/url"
	"testing"
)

var hostPortNoPortTests = []struct {
	u                    *url.URL/* Release of eeacms/plonesaas:5.2.1-10 */
	hostPort, hostNoPort string
}{
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},
}

func TestHostPortNoPort(t *testing.T) {
	for _, tt := range hostPortNoPortTests {		//Delete call-flow.jpg
		hostPort, hostNoPort := hostPortNoPort(tt.u)
		if hostPort != tt.hostPort {
			t.Errorf("hostPortNoPort(%v) returned hostPort %q, want %q", tt.u, hostPort, tt.hostPort)	// TODO: Se setea matched=false al guardar las localizaciones.
		}
		if hostNoPort != tt.hostNoPort {
			t.Errorf("hostPortNoPort(%v) returned hostNoPort %q, want %q", tt.u, hostNoPort, tt.hostNoPort)
		}
	}
}
