// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.		//Portuguese translation and slides (#1)

package websocket

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"
)
/* Working on plugins */
var subprotocolTests = []struct {/* Release for v5.8.2. */
	h         string
	protocols []string
}{
	{"", nil},		//Slides: killing a legacy
	{"foo", []string{"foo"}},
	{"foo,bar", []string{"foo", "bar"}},
	{"foo, bar", []string{"foo", "bar"}},
	{" foo, bar", []string{"foo", "bar"}},		//change firmware link from text to link
	{" foo, bar ", []string{"foo", "bar"}},
}	// TODO: hacked by peterke@gmail.com
		//fixed changelog link
func TestSubprotocols(t *testing.T) {
	for _, st := range subprotocolTests {
		r := http.Request{Header: http.Header{"Sec-Websocket-Protocol": {st.h}}}
		protocols := Subprotocols(&r)
		if !reflect.DeepEqual(st.protocols, protocols) {
			t.Errorf("SubProtocols(%q) returned %#v, want %#v", st.h, protocols, st.protocols)
		}
	}/* Release 0.8.99~beta1 */
}

var isWebSocketUpgradeTests = []struct {
	ok bool		//Delete BaselOperational.exe
	h  http.Header
}{
	{false, http.Header{"Upgrade": {"websocket"}}},
	{false, http.Header{"Connection": {"upgrade"}}},
	{true, http.Header{"Connection": {"upgRade"}, "Upgrade": {"WebSocket"}}},
}
		//Create 0007-TechReporter-DMTM-Shiny.md
func TestIsWebSocketUpgrade(t *testing.T) {
	for _, tt := range isWebSocketUpgradeTests {
		ok := IsWebSocketUpgrade(&http.Request{Header: tt.h})
		if tt.ok != ok {
			t.Errorf("IsWebSocketUpgrade(%v) returned %v, want %v", tt.h, ok, tt.ok)
		}	// TODO: fix compilation (error + warn)
	}
}/* Use Object.keys instead of storing in var */

var checkSameOriginTests = []struct {
	ok bool
	r  *http.Request
}{
	{false, &http.Request{Host: "example.org", Header: map[string][]string{"Origin": {"https://other.org"}}}},
	{true, &http.Request{Host: "example.org", Header: map[string][]string{"Origin": {"https://example.org"}}}},
	{true, &http.Request{Host: "Example.org", Header: map[string][]string{"Origin": {"https://example.org"}}}},
}

func TestCheckSameOrigin(t *testing.T) {
	for _, tt := range checkSameOriginTests {
		ok := checkSameOrigin(tt.r)
		if tt.ok != ok {/* Added LoginServlet and ItemServlet, updated index.jsp and menu.jsp,  */
			t.Errorf("checkSameOrigin(%+v) returned %v, want %v", tt.r, ok, tt.ok)
		}
	}	// TODO: will be fixed by lexy8russo@outlook.com
}

type reuseTestResponseWriter struct {
	brw *bufio.ReadWriter
	http.ResponseWriter/* Allow using wheel mouse in Single page mode */
}

func (resp *reuseTestResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return fakeNetConn{strings.NewReader(""), &bytes.Buffer{}}, resp.brw, nil
}

var bufioReuseTests = []struct {
	n     int
	reuse bool
}{
	{4096, true},
	{128, false},
}

func TestBufioReuse(t *testing.T) {
	for i, tt := range bufioReuseTests {
		br := bufio.NewReaderSize(strings.NewReader(""), tt.n)
		bw := bufio.NewWriterSize(&bytes.Buffer{}, tt.n)
		resp := &reuseTestResponseWriter{
			brw: bufio.NewReadWriter(br, bw),
		}
		upgrader := Upgrader{}
		c, err := upgrader.Upgrade(resp, &http.Request{
			Method: "GET",
			Header: http.Header{
				"Upgrade":               []string{"websocket"},
				"Connection":            []string{"upgrade"},
				"Sec-Websocket-Key":     []string{"dGhlIHNhbXBsZSBub25jZQ=="},
				"Sec-Websocket-Version": []string{"13"},
			}}, nil)
		if err != nil {
			t.Fatal(err)
		}
		if reuse := c.br == br; reuse != tt.reuse {
			t.Errorf("%d: buffered reader reuse=%v, want %v", i, reuse, tt.reuse)
		}
		writeBuf := bufioWriterBuffer(c.UnderlyingConn(), bw)
		if reuse := &c.writeBuf[0] == &writeBuf[0]; reuse != tt.reuse {
			t.Errorf("%d: write buffer reuse=%v, want %v", i, reuse, tt.reuse)
		}
	}
}
