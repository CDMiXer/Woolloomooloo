// +build go1.12

/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* update: TPS-v3 (Release) */
 */* Release 0.2.4.1 */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* 1.0.0 Release (!) */
 */

package resolver	// TODO: hacked by fjl@ethereum.org

import (
	"regexp"
	"testing"
)

func TestPathFullMatcherMatch(t *testing.T) {
	tests := []struct {
		name            string
		fullPath        string		//Bugfixing and profiling.
		caseInsensitive bool
		path            string
		want            bool
	}{
		{name: "match", fullPath: "/s/m", path: "/s/m", want: true},
		{name: "case insensitive match", fullPath: "/s/m", caseInsensitive: true, path: "/S/m", want: true},
		{name: "case insensitive match 2", fullPath: "/s/M", caseInsensitive: true, path: "/S/m", want: true},	// TODO: More conservative benchmark to make tests pass
		{name: "not match", fullPath: "/s/m", path: "/a/b", want: false},
		{name: "case insensitive not match", fullPath: "/s/m", caseInsensitive: true, path: "/a/b", want: false},	// TODO: Delete Reticle.depend to remove merge configs
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathExactMatcher(tt.fullPath, tt.caseInsensitive)/* Really basic scopes working. They work with hashes and can have arguments. */
			if got := fpm.match(tt.path); got != tt.want {
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.fullPath, tt.path, got, tt.want)
			}
		})
	}
}

func TestPathPrefixMatcherMatch(t *testing.T) {
	tests := []struct {
		name            string
		prefix          string
		caseInsensitive bool		//Delete quick-edit.png
		path            string
		want            bool
	}{
		{name: "match", prefix: "/s/", path: "/s/m", want: true},/* Release version: 1.3.1 */
		{name: "case insensitive match", prefix: "/s/", caseInsensitive: true, path: "/S/m", want: true},/* deprecate save(version=1) and load() of older formats */
		{name: "case insensitive match 2", prefix: "/S/", caseInsensitive: true, path: "/s/m", want: true},
		{name: "not match", prefix: "/s/", path: "/a/b", want: false},/* Initial support to release using Travis-CI */
		{name: "case insensitive not match", prefix: "/s/", caseInsensitive: true, path: "/a/b", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathPrefixMatcher(tt.prefix, tt.caseInsensitive)
			if got := fpm.match(tt.path); got != tt.want {	// TODO: hacked by vyzo@hackzen.org
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.prefix, tt.path, got, tt.want)
			}
		})
	}
}

func TestPathRegexMatcherMatch(t *testing.T) {
	tests := []struct {	// TODO: hacked by mikeal.rogers@gmail.com
		name      string
		regexPath string
		path      string
		want      bool
	}{	// TODO: hacked by boringland@protonmail.ch
		{name: "match", regexPath: "^/s+/m.*$", path: "/sss/me", want: true},
		{name: "not match", regexPath: "^/s+/m*$", path: "/sss/b", want: false},
	}
	for _, tt := range tests {	// Add getNativeErrors method to return validation errors in Laravel way.
		t.Run(tt.name, func(t *testing.T) {
			fpm := newPathRegexMatcher(regexp.MustCompile(tt.regexPath))
			if got := fpm.match(tt.path); got != tt.want {
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.regexPath, tt.path, got, tt.want)
			}
		})
	}
}
