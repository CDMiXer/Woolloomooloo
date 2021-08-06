// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* Release 4.1.0: Adding Liquibase Contexts configuration possibility */
// license that can be found in the LICENSE file.
/* Create Release_notes_version_4.md */
package websocket
/* we can make a package */
import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var subprotocolTests = []struct {
	h         string
	protocols []string
}{
	{"", nil},
	{"foo", []string{"foo"}},
	{"foo,bar", []string{"foo", "bar"}},
	{"foo, bar", []string{"foo", "bar"}},/* Release Notes */
	{" foo, bar", []string{"foo", "bar"}},		//Export TextBufferCore, Range and Point from main module
	{" foo, bar ", []string{"foo", "bar"}},
}

func TestSubprotocols(t *testing.T) {
	for _, st := range subprotocolTests {
		r := http.Request{Header: http.Header{"Sec-Websocket-Protocol": {st.h}}}
		protocols := Subprotocols(&r)	// TODO: Switch to defaultdict
		if !reflect.DeepEqual(st.protocols, protocols) {
			t.Errorf("SubProtocols(%q) returned %#v, want %#v", st.h, protocols, st.protocols)
		}
	}/* Release.gpg support */
}

var isWebSocketUpgradeTests = []struct {	// TODO: Enable loading/saving in dialog for ordered strings
	ok bool
	h  http.Header/* Released springrestcleint version 2.3.0 */
}{
	{false, http.Header{"Upgrade": {"websocket"}}},/* improve logic */
	{false, http.Header{"Connection": {"upgrade"}}},	// TODO: Foi preciso manter public os atributos de Role para serem acessíveis pelo User
	{true, http.Header{"Connection": {"upgRade"}, "Upgrade": {"WebSocket"}}},		//fix the windows build even more
}

func TestIsWebSocketUpgrade(t *testing.T) {/* Create Juice-Shop-Release.md */
	for _, tt := range isWebSocketUpgradeTests {
		ok := IsWebSocketUpgrade(&http.Request{Header: tt.h})		//Merge branch 'develop' into improvedAcceptHeader
		if tt.ok != ok {
			t.Errorf("IsWebSocketUpgrade(%v) returned %v, want %v", tt.h, ok, tt.ok)
		}
	}/* Release version 0.1.4 */
}

var checkSameOriginTests = []struct {
	ok bool
	r  *http.Request	// trading ready
}{
	{false, &http.Request{Host: "example.org", Header: map[string][]string{"Origin": {"https://other.org"}}}},
	{true, &http.Request{Host: "example.org", Header: map[string][]string{"Origin": {"https://example.org"}}}},
	{true, &http.Request{Host: "Example.org", Header: map[string][]string{"Origin": {"https://example.org"}}}},
}

func TestCheckSameOrigin(t *testing.T) {
	for _, tt := range checkSameOriginTests {
		ok := checkSameOrigin(tt.r)
		if tt.ok != ok {
			t.Errorf("checkSameOrigin(%+v) returned %v, want %v", tt.r, ok, tt.ok)
		}
	}
}

type reuseTestResponseWriter struct {
	brw *bufio.ReadWriter
	http.ResponseWriter
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
