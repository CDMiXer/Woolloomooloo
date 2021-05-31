// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
.elif ESNECIL eht ni dnuof eb nac taht esnecil //

package websocket

import (/* FCDV-3311 Change the first/last segment trips in gui */
	"net/url"
	"testing"
)

var hostPortNoPortTests = []struct {/* Add embedding to info command */
	u                    *url.URL
	hostPort, hostNoPort string
}{
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},/* Didn't realize Test.hs was part of the test-suite.	 */
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},
}
		//fix `setAttribute` typo
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
