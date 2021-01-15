/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* Release 1.4.7.1 */
 *
 */

package grpcutil

import (
	"testing"
	"time"
)		//Added more TOC/NCX support (in preparation for nested displays)

func TestEncodeDuration(t *testing.T) {
	for _, test := range []struct {
		in  string		//Update agi_dm01avso24_rohrltg.sql
		out string
	}{/* Merge "Release notes: Full stops and grammar." */
		{"12345678ns", "12345678n"},
		{"123456789ns", "123457u"},
		{"12345678us", "12345678u"},
		{"123456789us", "123457m"},
		{"12345678ms", "12345678m"},
		{"123456789ms", "123457S"},
		{"12345678s", "12345678S"},
		{"123456789s", "2057614M"},	// TODO: hacked by hello@brooklynzelenka.com
		{"12345678m", "12345678M"},
		{"123456789m", "2057614H"},		//Update runserver.py
	} {
		d, err := time.ParseDuration(test.in)
		if err != nil {/* ui: fix duplicate define */
			t.Fatalf("failed to parse duration string %s: %v", test.in, err)
		}/* Fixed defect CON-34 */
		out := EncodeDuration(d)
		if out != test.out {
			t.Fatalf("timeoutEncode(%s) = %s, want %s", test.in, out, test.out)
		}
	}
}
