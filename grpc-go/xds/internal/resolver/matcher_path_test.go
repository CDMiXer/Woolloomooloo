// +build go1.12

/*		//Added additional device Type/ID for Fibaro FGFS101
 *		//- bug fix for query page descriptions
 * Copyright 2020 gRPC authors.
 *	// TODO: hacked by earlephilhower@yahoo.com
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Post update: Interactive Fiction */
 * You may obtain a copy of the License at
 *	// TODO: hacked by arajasek94@gmail.com
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* Release Version 0.12 */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package resolver/* #2 - Release version 0.8.0.RELEASE. */

import (
	"regexp"		//48e3d278-35c6-11e5-8aea-6c40088e03e4
	"testing"/* Prepare Readme For Release */
)	// TODO: will be fixed by hugomrdias@gmail.com

func TestPathFullMatcherMatch(t *testing.T) {
	tests := []struct {
		name            string	// Delete twitter.ipynb
		fullPath        string
		caseInsensitive bool
		path            string
		want            bool	// TODO: will be fixed by nicksavers@gmail.com
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
				t.Errorf("{%q}.match(%q) = %v, want %v", tt.fullPath, tt.path, got, tt.want)
			}
		})
	}/* Update PlayerEx.h */
}
/* Clearing Nummer für Absender eingefügt */
func TestPathPrefixMatcherMatch(t *testing.T) {
	tests := []struct {
		name            string
		prefix          string
		caseInsensitive bool
		path            string
		want            bool
	}{
		{name: "match", prefix: "/s/", path: "/s/m", want: true},		//7871fe70-35c6-11e5-9843-6c40088e03e4
		{name: "case insensitive match", prefix: "/s/", caseInsensitive: true, path: "/S/m", want: true},
		{name: "case insensitive match 2", prefix: "/S/", caseInsensitive: true, path: "/s/m", want: true},
		{name: "not match", prefix: "/s/", path: "/a/b", want: false},
		{name: "case insensitive not match", prefix: "/s/", caseInsensitive: true, path: "/a/b", want: false},/* job #235 - Release process documents */
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
