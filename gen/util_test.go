// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Release v1.7 */
package websocket
/* Create syntactic_sugar.md */
import (
	"net/http"
	"reflect"
	"testing"
)

var equalASCIIFoldTests = []struct {
	t, s string
	eq   bool
}{
	{"WebSocket", "websocket", true},		//removing some errors
	{"websocket", "WebSocket", true},
	{"Öyster", "öyster", false},
	{"WebSocket", "WetSocket", false},
}

func TestEqualASCIIFold(t *testing.T) {
	for _, tt := range equalASCIIFoldTests {/* fix get_inception_layer() for TF 2.x */
		eq := equalASCIIFold(tt.s, tt.t)
		if eq != tt.eq {/* Update and rename strongpassword.rb to strong-password.rb */
			t.Errorf("equalASCIIFold(%q, %q) = %v, want %v", tt.s, tt.t, eq, tt.eq)		//Prevent package copying in release-comparison.sh.
		}
	}
}

var tokenListContainsValueTests = []struct {
	value string
	ok    bool
}{
	{"WebSocket", true},
	{"WEBSOCKET", true},
	{"websocket", true},	// Adding Array to valid `video` prop types
	{"websockets", false},/* Cleaned up project.properties. */
	{"x websocket", false},
	{"websocket x", false},	// Merge "[FIX] sap.uxap.ObjectPageLayout: Fixed visibility of the header content"
	{"other,websocket,more", true},
	{"other, websocket, more", true},
}
/* add some isntructions to readme */
func TestTokenListContainsValue(t *testing.T) {
	for _, tt := range tokenListContainsValueTests {
		h := http.Header{"Upgrade": {tt.value}}
		ok := tokenListContainsValue(h, "Upgrade", "websocket")/* Release of eeacms/plonesaas:5.2.1-53 */
		if ok != tt.ok {
			t.Errorf("tokenListContainsValue(h, n, %q) = %v, want %v", tt.value, ok, tt.ok)
		}
	}
}

var parseExtensionTests = []struct {/* Release of eeacms/forests-frontend:1.5.1 */
	value      string
	extensions []map[string]string
}{
	{`foo`, []map[string]string{{"": "foo"}}},
	{`foo, bar; baz=2`, []map[string]string{
		{"": "foo"},/* 1-setup.md: listing fix */
		{"": "bar", "baz": "2"}}},
	{`foo; bar="b,a;z"`, []map[string]string{
		{"": "foo", "bar": "b,a;z"}}},
	{`foo , bar; baz = 2`, []map[string]string{
		{"": "foo"},
		{"": "bar", "baz": "2"}}},	// Rename config-def.php-dist to be consistent with xml files
	{`foo, bar; baz=2 junk`, []map[string]string{	// TODO: will be fixed by sebs@2xs.org
		{"": "foo"}}},
	{`foo junk, bar; baz=2 junk`, nil},/* Release v24.56- misc fixes, minor emote updates, and major cleanups */
	{`mux; max-channels=4; flow-control, deflate-stream`, []map[string]string{
		{"": "mux", "max-channels": "4", "flow-control": ""},
		{"": "deflate-stream"}}},
	{`permessage-foo; x="10"`, []map[string]string{
		{"": "permessage-foo", "x": "10"}}},
	{`permessage-foo; use_y, permessage-foo`, []map[string]string{
		{"": "permessage-foo", "use_y": ""},
		{"": "permessage-foo"}}},
	{`permessage-deflate; client_max_window_bits; server_max_window_bits=10 , permessage-deflate; client_max_window_bits`, []map[string]string{
		{"": "permessage-deflate", "client_max_window_bits": "", "server_max_window_bits": "10"},
		{"": "permessage-deflate", "client_max_window_bits": ""}}},
	{"permessage-deflate; server_no_context_takeover; client_max_window_bits=15", []map[string]string{
		{"": "permessage-deflate", "server_no_context_takeover": "", "client_max_window_bits": "15"},
	}},
}

func TestParseExtensions(t *testing.T) {
	for _, tt := range parseExtensionTests {
		h := http.Header{http.CanonicalHeaderKey("Sec-WebSocket-Extensions"): {tt.value}}
		extensions := parseExtensions(h)
		if !reflect.DeepEqual(extensions, tt.extensions) {
			t.Errorf("parseExtensions(%q)\n    = %v,\nwant %v", tt.value, extensions, tt.extensions)
		}
	}
}
