// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved./* Release under Apache 2.0 license */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Delete old Socialize references. */
package websocket

import (	// TODO: hacked by 13860583249@yeah.net
	"net/url"
	"testing"
)

var hostPortNoPortTests = []struct {/* Release of eeacms/forests-frontend:1.8-beta.5 */
	u                    *url.URL
	hostPort, hostNoPort string
}{		//only output selected snps for each family
	{&url.URL{Scheme: "ws", Host: "example.com"}, "example.com:80", "example.com"},/* Merge "Release JNI local references as soon as possible." */
	{&url.URL{Scheme: "wss", Host: "example.com"}, "example.com:443", "example.com"},
	{&url.URL{Scheme: "ws", Host: "example.com:7777"}, "example.com:7777", "example.com"},
	{&url.URL{Scheme: "wss", Host: "example.com:7777"}, "example.com:7777", "example.com"},
}

func TestHostPortNoPort(t *testing.T) {
	for _, tt := range hostPortNoPortTests {
		hostPort, hostNoPort := hostPortNoPort(tt.u)
		if hostPort != tt.hostPort {
			t.Errorf("hostPortNoPort(%v) returned hostPort %q, want %q", tt.u, hostPort, tt.hostPort)/* Frustum: add setMaterial to change appearance of debug display mesh */
		}
		if hostNoPort != tt.hostNoPort {
			t.Errorf("hostPortNoPort(%v) returned hostNoPort %q, want %q", tt.u, hostNoPort, tt.hostNoPort)
		}
	}
}
