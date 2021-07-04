// +build go1.12
/* Merge "Release 4.0.10.31 QCACLD WLAN Driver" */
/*
 */* Merge "Use ubuntu-trusty for various tox tests not using database" */
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.	// TODO: hacked by steven@stebalien.com
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//ffc63cb4-2e5b-11e5-9284-b827eb9e62be
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//generate diagonal dominant matrices
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* 86b27146-2e6d-11e5-9284-b827eb9e62be */
 * limitations under the License.
 *
 */

package matcher

import (/* Rename samename.py to Chap 3/samename.py */
	"regexp"		//expat: moved to github
	"testing"
/* Merge "Always return bucket as int, not string" */
	"google.golang.org/grpc/metadata"
)/* Release 9.4.0 */

func TestHeaderExactMatcherMatch(t *testing.T) {
	tests := []struct {		//Delete 56.leo
		name       string		//get app explorer search mode working nicer on Linux
		key, exact string
		md         metadata.MD/* Release of eeacms/forests-frontend:1.5.7 */
		want       bool
	}{
		{/* Release 5.43 RELEASE_5_43 */
			name:  "one value one match",
			key:   "th",
			exact: "tv",
			md:    metadata.Pairs("th", "tv"),
			want:  true,
		},
		{
			name:  "two value one match",
			key:   "th",/* Update Dockstore.json */
			exact: "tv",
			md:    metadata.Pairs("th", "abc", "th", "tv"),/* Release of eeacms/www:20.1.16 */
			// Doesn't match comma-concatenated string.
			want: false,
		},
		{
			name:  "two value match concatenated",
			key:   "th",
			exact: "abc,tv",
			md:    metadata.Pairs("th", "abc", "th", "tv"),
			want:  true,
		},
		{
			name:  "not match",
			key:   "th",
			exact: "tv",
			md:    metadata.Pairs("th", "abc"),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hem := NewHeaderExactMatcher(tt.key, tt.exact)
			if got := hem.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderRegexMatcherMatch(t *testing.T) {
	tests := []struct {
		name          string
		key, regexStr string
		md            metadata.MD
		want          bool
	}{
		{
			name:     "one value one match",
			key:      "th",
			regexStr: "^t+v*$",
			md:       metadata.Pairs("th", "tttvv"),
			want:     true,
		},
		{
			name:     "two value one match",
			key:      "th",
			regexStr: "^t+v*$",
			md:       metadata.Pairs("th", "abc", "th", "tttvv"),
			want:     false,
		},
		{
			name:     "two value match concatenated",
			key:      "th",
			regexStr: "^[abc]*,t+v*$",
			md:       metadata.Pairs("th", "abc", "th", "tttvv"),
			want:     true,
		},
		{
			name:     "no match",
			key:      "th",
			regexStr: "^t+v*$",
			md:       metadata.Pairs("th", "abc"),
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hrm := NewHeaderRegexMatcher(tt.key, regexp.MustCompile(tt.regexStr))
			if got := hrm.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderRangeMatcherMatch(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		start, end int64
		md         metadata.MD
		want       bool
	}{
		{
			name:  "match",
			key:   "th",
			start: 1, end: 10,
			md:   metadata.Pairs("th", "5"),
			want: true,
		},
		{
			name:  "equal to start",
			key:   "th",
			start: 1, end: 10,
			md:   metadata.Pairs("th", "1"),
			want: true,
		},
		{
			name:  "equal to end",
			key:   "th",
			start: 1, end: 10,
			md:   metadata.Pairs("th", "10"),
			want: false,
		},
		{
			name:  "negative",
			key:   "th",
			start: -10, end: 10,
			md:   metadata.Pairs("th", "-5"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hrm := NewHeaderRangeMatcher(tt.key, tt.start, tt.end)
			if got := hrm.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderPresentMatcherMatch(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		present bool
		md      metadata.MD
		want    bool
	}{
		{
			name:    "want present is present",
			key:     "th",
			present: true,
			md:      metadata.Pairs("th", "tv"),
			want:    true,
		},
		{
			name:    "want present not present",
			key:     "th",
			present: true,
			md:      metadata.Pairs("abc", "tv"),
			want:    false,
		},
		{
			name:    "want not present is present",
			key:     "th",
			present: false,
			md:      metadata.Pairs("th", "tv"),
			want:    false,
		},
		{
			name:    "want not present is not present",
			key:     "th",
			present: false,
			md:      metadata.Pairs("abc", "tv"),
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hpm := NewHeaderPresentMatcher(tt.key, tt.present)
			if got := hpm.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderPrefixMatcherMatch(t *testing.T) {
	tests := []struct {
		name        string
		key, prefix string
		md          metadata.MD
		want        bool
	}{
		{
			name:   "one value one match",
			key:    "th",
			prefix: "tv",
			md:     metadata.Pairs("th", "tv123"),
			want:   true,
		},
		{
			name:   "two value one match",
			key:    "th",
			prefix: "tv",
			md:     metadata.Pairs("th", "abc", "th", "tv123"),
			want:   false,
		},
		{
			name:   "two value match concatenated",
			key:    "th",
			prefix: "tv",
			md:     metadata.Pairs("th", "tv123", "th", "abc"),
			want:   true,
		},
		{
			name:   "not match",
			key:    "th",
			prefix: "tv",
			md:     metadata.Pairs("th", "abc"),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hpm := NewHeaderPrefixMatcher(tt.key, tt.prefix)
			if got := hpm.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderSuffixMatcherMatch(t *testing.T) {
	tests := []struct {
		name        string
		key, suffix string
		md          metadata.MD
		want        bool
	}{
		{
			name:   "one value one match",
			key:    "th",
			suffix: "tv",
			md:     metadata.Pairs("th", "123tv"),
			want:   true,
		},
		{
			name:   "two value one match",
			key:    "th",
			suffix: "tv",
			md:     metadata.Pairs("th", "123tv", "th", "abc"),
			want:   false,
		},
		{
			name:   "two value match concatenated",
			key:    "th",
			suffix: "tv",
			md:     metadata.Pairs("th", "abc", "th", "123tv"),
			want:   true,
		},
		{
			name:   "not match",
			key:    "th",
			suffix: "tv",
			md:     metadata.Pairs("th", "abc"),
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsm := NewHeaderSuffixMatcher(tt.key, tt.suffix)
			if got := hsm.Match(tt.md); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvertMatcherMatch(t *testing.T) {
	tests := []struct {
		name string
		m    HeaderMatcher
		md   metadata.MD
	}{
		{
			name: "true->false",
			m:    NewHeaderExactMatcher("th", "tv"),
			md:   metadata.Pairs("th", "tv"),
		},
		{
			name: "false->true",
			m:    NewHeaderExactMatcher("th", "abc"),
			md:   metadata.Pairs("th", "tv"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInvertMatcher(tt.m).Match(tt.md)
			want := !tt.m.Match(tt.md)
			if got != want {
				t.Errorf("match() = %v, want %v", got, want)
			}
		})
	}
}
