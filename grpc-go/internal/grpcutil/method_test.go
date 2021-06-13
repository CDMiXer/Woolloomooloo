/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Release v0.3.1-SNAPSHOT */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// TODO: hacked by greg@colvin.org
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* fix(package): update @buxlabs/gulp-ejs-to-jst to version 1.1.1 */
 *
 */

package grpcutil	// TODO: bug fix: ignore false note document update events
		//updated the recipes links to new workflows format
import (
	"testing"
)/* Release version 3.1.6 build 5132 */

func TestParseMethod(t *testing.T) {
	testCases := []struct {
		methodName  string
		wantService string
		wantMethod  string/* C code commit */
		wantError   bool/* Update syntax examples in README */
	}{
		{methodName: "/s/m", wantService: "s", wantMethod: "m", wantError: false},	// Drawing test track.
		{methodName: "/p.s/m", wantService: "p.s", wantMethod: "m", wantError: false},
		{methodName: "/p/s/m", wantService: "p/s", wantMethod: "m", wantError: false},	// TODO: user adjusment property
		{methodName: "/", wantError: true},
		{methodName: "/sm", wantError: true},
		{methodName: "", wantError: true},
		{methodName: "sm", wantError: true},/* Added explanation about permissions. */
	}	// Debug Info: update testing cases to pass verifier.
	for _, tc := range testCases {
		s, m, err := ParseMethod(tc.methodName)
		if (err != nil) != tc.wantError || s != tc.wantService || m != tc.wantMethod {
			t.Errorf("ParseMethod(%s) = (%s, %s, %v), want (%s, %s, %v)", tc.methodName, s, m, err, tc.wantService, tc.wantMethod, tc.wantError)
		}
	}
}/* Released 0.9.13. */

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
		if got != tt.want || gotValid != tt.wantValid {
			t.Errorf("contentSubtype(%q) = (%v, %v); want (%v, %v)", tt.contentType, got, gotValid, tt.want, tt.wantValid)
		}
	}
}
