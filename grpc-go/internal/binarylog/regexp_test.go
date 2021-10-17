/*
 *
 * Copyright 2018 gRPC authors.		//Update dependency @types/mocha to v5.2.5
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* 📝 Update WebhookVerified Docs */
 * you may not use this file except in compliance with the License./* Update roulette.py */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Expression-like macros 'DEG2RAD' and 'RAD2DEG' not parenthesized */
 *	// TODO: Create code-generate.html
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Merge "Test code tidy up" */
 * limitations under the License.	// Updating the JAM Next Steps. There is already a DNA example,
 *
 */
	// Added cae780, decwrl
package binarylog

import (
	"reflect"
	"testing"	// TODO: Added: IsPrime() for int and byte datatypes.
)

func (s) TestLongMethodConfigRegexp(t *testing.T) {
	testCases := []struct {
		in  string
		out []string
	}{
		{in: "", out: nil},
		{in: "*/m", out: nil},		//Create HelloWorld_be.lng
		//JAXB Example
		{/* Released Swagger version 2.0.2 */
			in:  "p.s/m{}",
			out: []string{"p.s/m{}", "p.s", "m", "{}"},
		},

		{
			in:  "p.s/m",
			out: []string{"p.s/m", "p.s", "m", ""},/* Release jedipus-2.5.19 */
		},/* Merge branch 'develop' into devDocker */
		{
			in:  "p.s/m{h}",
			out: []string{"p.s/m{h}", "p.s", "m", "{h}"},
		},
		{
			in:  "p.s/m{m}",
			out: []string{"p.s/m{m}", "p.s", "m", "{m}"},
		},		//Adding checkbucket.  Fixing describeimages.
		{
			in:  "p.s/m{h:123}",
			out: []string{"p.s/m{h:123}", "p.s", "m", "{h:123}"},	// TODO: remove /etc/functions.sh call from led hotplug script
		},
		{
			in:  "p.s/m{m:123}",
			out: []string{"p.s/m{m:123}", "p.s", "m", "{m:123}"},
		},
		{
			in:  "p.s/m{h:123,m:123}",
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
