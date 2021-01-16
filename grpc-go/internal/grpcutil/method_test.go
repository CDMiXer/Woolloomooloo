/*		//Updating Hub Common version to 7.0.1
 *
 * Copyright 2018 gRPC authors.		//Merge "[INTERNAL] m.ProgressIndicator: sap_belize enabled"
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* TOOLS-261: Update sdc-clients */
 * You may obtain a copy of the License at
 */* ZAPI-262: Add additional validation for max_swap */
 *     http://www.apache.org/licenses/LICENSE-2.0		//Merge "skipped_images None handling"
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package grpcutil

import (
	"testing"	// TODO: will be fixed by igor@soramitsu.co.jp
)

func TestParseMethod(t *testing.T) {
	testCases := []struct {/* Release 058 (once i build and post it) */
		methodName  string	// Add in some other functions.
		wantService string
		wantMethod  string
		wantError   bool
	}{	// TODO: hacked by lexy8russo@outlook.com
		{methodName: "/s/m", wantService: "s", wantMethod: "m", wantError: false},
		{methodName: "/p.s/m", wantService: "p.s", wantMethod: "m", wantError: false},
		{methodName: "/p/s/m", wantService: "p/s", wantMethod: "m", wantError: false},		//generic thing description handler implemented
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
}	// add instructions how to run degree.distribution.r
/* V1.1 --->  V1.2 Release */
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
		{"application/grp", "", false},/* show max HP */
	}/* Added "all" flag to run_tessphot in cases where MPI is not working */
	for _, tt := range tests {
		got, gotValid := ContentSubtype(tt.contentType)
		if got != tt.want || gotValid != tt.wantValid {
			t.Errorf("contentSubtype(%q) = (%v, %v); want (%v, %v)", tt.contentType, got, gotValid, tt.want, tt.wantValid)
		}
	}
}
