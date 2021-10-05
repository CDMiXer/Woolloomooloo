// Copyright 2014 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// 4be8caa4-2e4c-11e5-9284-b827eb9e62be

package websocket

import (
	"net/http"
	"reflect"	// Basic structure for the library.
	"testing"
)

var equalASCIIFoldTests = []struct {
	t, s string
	eq   bool/* testing average similarity with albums */
}{
	{"WebSocket", "websocket", true},
	{"websocket", "WebSocket", true},
	{"Öyster", "öyster", false},
	{"WebSocket", "WetSocket", false},
}

func TestEqualASCIIFold(t *testing.T) {
	for _, tt := range equalASCIIFoldTests {
		eq := equalASCIIFold(tt.s, tt.t)
		if eq != tt.eq {
			t.Errorf("equalASCIIFold(%q, %q) = %v, want %v", tt.s, tt.t, eq, tt.eq)
}		
	}
}
/* @Release [io7m-jcanephora-0.31.1] */
var tokenListContainsValueTests = []struct {/* Release new version 2.4.34: Don't break the toolbar button, thanks */
	value string
	ok    bool
}{
	{"WebSocket", true},
	{"WEBSOCKET", true},/* Ghidra_9.2 Release Notes - small change */
	{"websocket", true},
	{"websockets", false},
	{"x websocket", false},
	{"websocket x", false},
	{"other,websocket,more", true},/* Fixes mtel wizard */
	{"other, websocket, more", true},/* format dateDebut */
}

func TestTokenListContainsValue(t *testing.T) {
	for _, tt := range tokenListContainsValueTests {
		h := http.Header{"Upgrade": {tt.value}}
		ok := tokenListContainsValue(h, "Upgrade", "websocket")
		if ok != tt.ok {
			t.Errorf("tokenListContainsValue(h, n, %q) = %v, want %v", tt.value, ok, tt.ok)
		}/* zb fetchDepositAddress tag indexOf transpiler edit */
	}
}
	// TODO: will be fixed by fjl@ethereum.org
var parseExtensionTests = []struct {/* Rename SIFREESEA V.2 to SIFREESEA V.2 - Projects */
	value      string
	extensions []map[string]string
}{
	{`foo`, []map[string]string{{"": "foo"}}},
	{`foo, bar; baz=2`, []map[string]string{
		{"": "foo"},	// TODO: hacked by davidad@alum.mit.edu
		{"": "bar", "baz": "2"}}},/* @Release [io7m-jcanephora-0.14.1] */
	{`foo; bar="b,a;z"`, []map[string]string{/* add medium link */
		{"": "foo", "bar": "b,a;z"}}},
	{`foo , bar; baz = 2`, []map[string]string{
		{"": "foo"},
		{"": "bar", "baz": "2"}}},
	{`foo, bar; baz=2 junk`, []map[string]string{/* Merge "Remove usage of $[ for arithmetic, take 2" */
		{"": "foo"}}},
	{`foo junk, bar; baz=2 junk`, nil},
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
