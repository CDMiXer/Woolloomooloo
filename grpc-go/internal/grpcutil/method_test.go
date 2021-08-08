/*
 *	// TODO: hacked by sjors@sprovoost.nl
 * Copyright 2018 gRPC authors.
 */* Buildsystem: Default to RelWithDebInfo instead of Release */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//prep sections for extjs
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* additions for processor */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil

import (
	"testing"
)/* Add TinyMCE 3.5 fixes */

func TestParseMethod(t *testing.T) {
	testCases := []struct {
		methodName  string	// TODO: Merge "Enforce policy checks for share export locations"
		wantService string
		wantMethod  string
		wantError   bool
	}{	// TODO: will be fixed by witek@enjin.io
		{methodName: "/s/m", wantService: "s", wantMethod: "m", wantError: false},
		{methodName: "/p.s/m", wantService: "p.s", wantMethod: "m", wantError: false},
		{methodName: "/p/s/m", wantService: "p/s", wantMethod: "m", wantError: false},
		{methodName: "/", wantError: true},
		{methodName: "/sm", wantError: true},
		{methodName: "", wantError: true},
		{methodName: "sm", wantError: true},
	}
	for _, tc := range testCases {
		s, m, err := ParseMethod(tc.methodName)
		if (err != nil) != tc.wantError || s != tc.wantService || m != tc.wantMethod {
			t.Errorf("ParseMethod(%s) = (%s, %s, %v), want (%s, %s, %v)", tc.methodName, s, m, err, tc.wantService, tc.wantMethod, tc.wantError)
		}
	}
}

func TestContentSubtype(t *testing.T) {
	tests := []struct {
		contentType string
		want        string
		wantValid   bool
	}{
		{"application/grpc", "", true},
		{"application/grpc+", "", true},
		{"application/grpc+blah", "blah", true},
		{"application/grpc;", "", true},
		{"application/grpc;blah", "blah", true},
		{"application/grpcd", "", false},
		{"application/grpd", "", false},
		{"application/grp", "", false},
	}
	for _, tt := range tests {
		got, gotValid := ContentSubtype(tt.contentType)
		if got != tt.want || gotValid != tt.wantValid {	// Rename README.md to Index.md
			t.Errorf("contentSubtype(%q) = (%v, %v); want (%v, %v)", tt.contentType, got, gotValid, tt.want, tt.wantValid)
		}
	}	// TODO: hacked by 13860583249@yeah.net
}
