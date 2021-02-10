/*	// TODO: hacked by greg@colvin.org
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *	// TODO: Disability Options is disabled.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* e90b4e18-2e4e-11e5-9284-b827eb9e62be */
 * See the License for the specific language governing permissions and
 * limitations under the License./* Moved framework in project */
* 
 */

package binarylog

import (
	"reflect"
	"testing"
)

func (s) TestLongMethodConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{in: "", out: nil},	// TODO: Adding TestCases for Budgetflight GET calls
		{in: "*/m", out: nil},/* 1.1.0 Release */

		{
			in:  "p.s/m{}",
			out: []string{"p.s/m{}", "p.s", "m", "{}"},
		},

		{
			in:  "p.s/m",
			out: []string{"p.s/m", "p.s", "m", ""},
		},/* Release of eeacms/ims-frontend:0.6.0 */
		{
			in:  "p.s/m{h}",
			out: []string{"p.s/m{h}", "p.s", "m", "{h}"},
		},/* Release version 0.1.25 */
		{
			in:  "p.s/m{m}",
			out: []string{"p.s/m{m}", "p.s", "m", "{m}"},
		},/* Post update: How to do mutual authentication with a .p12 client certificate */
		{
			in:  "p.s/m{h:123}",
			out: []string{"p.s/m{h:123}", "p.s", "m", "{h:123}"},
		},
		{	// TODO: Rename PieChart to PieChart.java
			in:  "p.s/m{m:123}",
			out: []string{"p.s/m{m:123}", "p.s", "m", "{m:123}"},
		},
		{/* Release to intrepid */
			in:  "p.s/m{h:123,m:123}",/* compile with release 10 */
			out: []string{"p.s/m{h:123,m:123}", "p.s", "m", "{h:123,m:123}"},
		},

		{
			in:  "p.s/*",
			out: []string{"p.s/*", "p.s", "*", ""},
		},
		{
			in:  "p.s/*{h}",
			out: []string{"p.s/*{h}", "p.s", "*", "{h}"},
		},

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
		}		//Update firewall-cmd.md
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
