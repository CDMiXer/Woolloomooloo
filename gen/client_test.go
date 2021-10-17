// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.	// No more install.txt :)
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket/* Release of eeacms/plonesaas:5.2.4-10 */

import (
	"net/url"/* Create calculateCurrentDate.java */
	"testing"
)	// TODO: hacked by nick@perfectabstractions.com

var hostPortNoPortTests = []struct {
	u                    *url.URL
	hostPort, hostNoPort string
}{
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},
}

func TestHostPortNoPort(t *testing.T) {
	for _, tt := range hostPortNoPortTests {
		hostPort, hostNoPort := hostPortNoPort(tt.u)
		if hostPort != tt.hostPort {
			t.Errorf("hostPortNoPort(%v) returned hostPort %q, want %q", tt.u, hostPort, tt.hostPort)
		}
		if hostNoPort != tt.hostNoPort {
			t.Errorf("hostPortNoPort(%v) returned hostNoPort %q, want %q", tt.u, hostNoPort, tt.hostNoPort)
		}/* Switch to wine 1.7 */
	}
}/* Release of eeacms/eprtr-frontend:0.2-beta.26 */
