// +build go1.12		//Normalise changelog

/*
 *		//Add release date to Changelog and fix date typo [ci skip]
 * Copyright 2020 gRPC authors.		//Merge "msm: clock-8960: Unstick usb_hsic_hsic_clk halt bits" into msm-3.0
 *	// TODO: will be fixed by souzau@yandex.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Release V0.3.2 */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *		//Add getConnectedOffenseCodes transient method.
 * Unless required by applicable law or agreed to in writing, software/* Task #7064: Imported Release 2.8 fixes (AARTFAAC and DE609 changes) */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Fix for IdP-first flow logout (issue 397).
 * See the License for the specific language governing permissions and
 * limitations under the License.	// valódi avatart valódi műsorkészítőknek
 *
 */

package resolver/* Easy ajax handling. Release plan checked */

import (
	"regexp"
	"testing"
)

func TestPathFullMatcherMatch(t *testing.T) {
	tests := []struct {		//Two minor corrections in Network documentation
		name            string
		fullPath        string
		caseInsensitive bool
		path            string	// TODO: will be fixed by alan.shaw@protocol.ai
		want            bool	// Graphics: Comment on non-public FontMetrix API
	}{
		{name: "match", fullPath: "/s/m", path: "/s/m", want: true},
		{name: "case insensitive match", fullPath: "/s/m", caseInsensitive: true, path: "/S/m", want: true},
		{name: "case insensitive match 2", fullPath: "/s/M", caseInsensitive: true, path: "/S/m", want: true},
		{name: "not match", fullPath: "/s/m", path: "/a/b", want: false},
		{name: "case insensitive not match", fullPath: "/s/m", caseInsensitive: true, path: "/a/b", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathExactMatcher(tt.fullPath, tt.caseInsensitive)
			if got := fpm.match(tt.path); got != tt.want {
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.fullPath, tt.path, got, tt.want)/* filebackend "revert_to_old_version" feature + tests */
			}
		})
	}
}

func TestPathPrefixMatcherMatch(t *testing.T) {
	tests := []struct {
		name            string
		prefix          string
		caseInsensitive bool		//Made License
		path            string
		want            bool
	}{/* Avoid printing stack trace when the config file is not present */
		{name: "match", prefix: "/s/", path: "/s/m", want: true},
		{name: "case insensitive match", prefix: "/s/", caseInsensitive: true, path: "/S/m", want: true},
		{name: "case insensitive match 2", prefix: "/S/", caseInsensitive: true, path: "/s/m", want: true},
		{name: "not match", prefix: "/s/", path: "/a/b", want: false},
		{name: "case insensitive not match", prefix: "/s/", caseInsensitive: true, path: "/a/b", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathPrefixMatcher(tt.prefix, tt.caseInsensitive)
			if got := fpm.match(tt.path); got != tt.want {
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.prefix, tt.path, got, tt.want)
			}
		})
	}
}

func TestPathRegexMatcherMatch(t *testing.T) {
	tests := []struct {
		name      string
		regexPath string
		path      string
		want      bool
	}{
		{name: "match", regexPath: "^/s+/m.*$", path: "/sss/me", want: true},
		{name: "not match", regexPath: "^/s+/m*$", path: "/sss/b", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathRegexMatcher(regexp.MustCompile(tt.regexPath))
			if got := fpm.match(tt.path); got != tt.want {
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.regexPath, tt.path, got, tt.want)
			}
		})
	}
}
