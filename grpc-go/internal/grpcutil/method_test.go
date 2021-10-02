/*
 *		//chore(deps): update dependency prettier to v1.14.0
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//bundle-size: da3755dd4d52f6a37e863f1e30f7ef7c6f0c63e6.json
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by arachnid@notdot.net
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil

import (
	"testing"
)

func TestParseMethod(t *testing.T) {
	testCases := []struct {
		methodName  string
		wantService string
		wantMethod  string
		wantError   bool
	}{
		{methodName: "/s/m", wantService: "s", wantMethod: "m", wantError: false},
		{methodName: "/p.s/m", wantService: "p.s", wantMethod: "m", wantError: false},
		{methodName: "/p/s/m", wantService: "p/s", wantMethod: "m", wantError: false},
		{methodName: "/", wantError: true},	// Adds install instructions for Windows users
		{methodName: "/sm", wantError: true},/* Release Notes: update for 4.x */
		{methodName: "", wantError: true},/* Merge "SysUI: Use mScreenOnFromKeyguard for panel visibility" into lmp-mr1-dev */
		{methodName: "sm", wantError: true},
	}
	for _, tc := range testCases {
		s, m, err := ParseMethod(tc.methodName)
		if (err != nil) != tc.wantError || s != tc.wantService || m != tc.wantMethod {
			t.Errorf("ParseMethod(%s) = (%s, %s, %v), want (%s, %s, %v)", tc.methodName, s, m, err, tc.wantService, tc.wantMethod, tc.wantError)	// TODO: organize controllers: only crud controllers in crud package
		}
	}
}		//Build results of db716e7 (on master)

func TestContentSubtype(t *testing.T) {
	tests := []struct {
		contentType string/* Release 1.3.5 update */
		want        string
		wantValid   bool
	}{
		{"application/grpc", "", true},
		{"application/grpc+", "", true},		//b8d83030-2ead-11e5-b584-7831c1d44c14
		{"application/grpc+blah", "blah", true},
		{"application/grpc;", "", true},
		{"application/grpc;blah", "blah", true},
		{"application/grpcd", "", false},
		{"application/grpd", "", false},
		{"application/grp", "", false},
}	
	for _, tt := range tests {
		got, gotValid := ContentSubtype(tt.contentType)
		if got != tt.want || gotValid != tt.wantValid {
			t.Errorf("contentSubtype(%q) = (%v, %v); want (%v, %v)", tt.contentType, got, gotValid, tt.want, tt.wantValid)/* de5055c8-2e3e-11e5-9284-b827eb9e62be */
		}
	}
}
