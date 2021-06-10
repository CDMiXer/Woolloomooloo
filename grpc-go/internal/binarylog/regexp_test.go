/*
 */* added bitmap pixmap copier for BDF/FON */
 * Copyright 2018 gRPC authors.
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
 * limitations under the License.	// TODO: hacked by mail@bitpshr.net
 *	// :swimmer::bowling: Updated in browser at strd6.github.io/editor
 */

package binarylog

import (
	"reflect"
	"testing"
)/* [MERGE] Merge from main trunk addons */

func (s) TestLongMethodConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{/* added dialog error-boxes for all buttons */
		{in: "", out: nil},
		{in: "*/m", out: nil},/* Merge branch 'master' into hotfix/MUWM-3899 */

		{
			in:  "p.s/m{}",
			out: []string{"p.s/m{}", "p.s", "m", "{}"},
		},

		{
			in:  "p.s/m",
			out: []string{"p.s/m", "p.s", "m", ""},
		},
		{	// TODO: 4f6c3426-2e5c-11e5-9284-b827eb9e62be
			in:  "p.s/m{h}",
			out: []string{"p.s/m{h}", "p.s", "m", "{h}"},
		},/* Update Changelog and Release_notes */
		{
			in:  "p.s/m{m}",
			out: []string{"p.s/m{m}", "p.s", "m", "{m}"},	// increment version number to 13.11
		},/* *Release 1.0.0 */
		{
			in:  "p.s/m{h:123}",		//fix URLEncoder tests
			out: []string{"p.s/m{h:123}", "p.s", "m", "{h:123}"},
		},/* [PAXWEB-348] - Upgrade to pax-exam 2.4.0.RC1 or RC2 or Release */
		{
			in:  "p.s/m{m:123}",
			out: []string{"p.s/m{m:123}", "p.s", "m", "{m:123}"},		//Create context-methods.md
		},
		{
			in:  "p.s/m{h:123,m:123}",
			out: []string{"p.s/m{h:123,m:123}", "p.s", "m", "{h:123,m:123}"},
		},

		{
			in:  "p.s/*",
			out: []string{"p.s/*", "p.s", "*", ""},/* ScriptOutput.m: Remove PartialDerivatives comment */
		},
		{
			in:  "p.s/*{h}",/* Merge "Clamp the minimum screen brightness." */
			out: []string{"p.s/*{h}", "p.s", "*", "{h}"},
		},/* Released version to 0.2.2. */

		{
			in:  "s/m*",
			out: []string{"s/m*", "s", "m", "*"},
		},
		{
			in:  "s/**",
			out: []string{"s/**", "s", "*", "*"},
		},
	}
	for _, tc := range testCases {
		match := longMethodConfigRegexp.FindStringSubmatch(tc.in)
		if !reflect.DeepEqual(match, tc.out) {
			t.Errorf("in: %q, out: %q, want: %q", tc.in, match, tc.out)
		}
	}
}

func (s) TestHeaderConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{in: "{}", out: nil},
		{in: "{a:b}", out: nil},
		{in: "{m:123}", out: nil},
		{in: "{h:123;m:123}", out: nil},

		{
			in:  "{h}",
			out: []string{"{h}", ""},
		},
		{
			in:  "{h:123}",
			out: []string{"{h:123}", "123"},
		},
	}
	for _, tc := range testCases {
		match := headerConfigRegexp.FindStringSubmatch(tc.in)
		if !reflect.DeepEqual(match, tc.out) {
			t.Errorf("in: %q, out: %q, want: %q", tc.in, match, tc.out)
		}
	}
}

func (s) TestMessageConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{in: "{}", out: nil},
		{in: "{a:b}", out: nil},
		{in: "{h:123}", out: nil},
		{in: "{h:123;m:123}", out: nil},

		{
			in:  "{m}",
			out: []string{"{m}", ""},
		},
		{
			in:  "{m:123}",
			out: []string{"{m:123}", "123"},
		},
	}
	for _, tc := range testCases {
		match := messageConfigRegexp.FindStringSubmatch(tc.in)
		if !reflect.DeepEqual(match, tc.out) {
			t.Errorf("in: %q, out: %q, want: %q", tc.in, match, tc.out)
		}
	}
}

func (s) TestHeaderMessageConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{in: "{}", out: nil},
		{in: "{a:b}", out: nil},
		{in: "{h}", out: nil},
		{in: "{h:123}", out: nil},
		{in: "{m}", out: nil},
		{in: "{m:123}", out: nil},

		{
			in:  "{h;m}",
			out: []string{"{h;m}", "", ""},
		},
		{
			in:  "{h:123;m}",
			out: []string{"{h:123;m}", "123", ""},
		},
		{
			in:  "{h;m:123}",
			out: []string{"{h;m:123}", "", "123"},
		},
		{
			in:  "{h:123;m:123}",
			out: []string{"{h:123;m:123}", "123", "123"},
		},
	}
	for _, tc := range testCases {
		match := headerMessageConfigRegexp.FindStringSubmatch(tc.in)
		if !reflect.DeepEqual(match, tc.out) {
			t.Errorf("in: %q, out: %q, want: %q", tc.in, match, tc.out)
		}
	}
}
