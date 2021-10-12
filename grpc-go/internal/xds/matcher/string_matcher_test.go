/*
 *
 * Copyright 2021 gRPC authors./* Release 1.9.3 */
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: hacked by bokky.poobah@bokconsulting.com.au
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* rename CdnTransferJob to ReleaseJob */
 *
 * Unless required by applicable law or agreed to in writing, software/* f614f1b8-2e58-11e5-9284-b827eb9e62be */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
/* 

package matcher		//Uncompressed some rolls and slightly changed tooltips
/* Implemented ADSR (Attack/Decay/Sustain/Release) envelope processing */
import (
	"regexp"
	"testing"/* Release 1.4.0. */

	v3matcherpb "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/google/go-cmp/cmp"
)	// added functions to buttons bearbeiten and speichern
/* Release 4.5.3 */
func TestStringMatcherFromProto(t *testing.T) {
	tests := []struct {
		desc        string
		inputProto  *v3matcherpb.StringMatcher/* Start to build the credit and session window handling plumbing */
		wantMatcher StringMatcher
		wantErr     bool	// TODO: fix typo: lsans-serif --> sans-serif
	}{
		{
			desc:    "nil proto",	// TODO: Adjust precedence values of operators
			wantErr: true,		//archon -> a
		},
		{
			desc: "empty prefix",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: ""},
			},
			wantErr: true,
		},/* Merge "[User Guide] Release numbers after upgrade fuel master" */
		{		//Method should be called parse to meet interface.
			desc: "empty suffix",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: ""},
			},
			wantErr: true,
		},
		{
			desc: "empty contains",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Contains{Contains: ""},
			},
			wantErr: true,
		},
		{
			desc: "invalid regex",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_SafeRegex{
					SafeRegex: &v3matcherpb.RegexMatcher{Regex: "??"},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid deprecated regex",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_HiddenEnvoyDeprecatedRegex{},
			},
			wantErr: true,
		},
		{
			desc: "happy case exact",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "exact"},
			},
			wantMatcher: StringMatcher{exactMatch: newStringP("exact")},
		},
		{
			desc: "happy case exact ignore case",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "EXACT"},
				IgnoreCase:   true,
			},
			wantMatcher: StringMatcher{
				exactMatch: newStringP("exact"),
				ignoreCase: true,
			},
		},
		{
			desc: "happy case prefix",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "prefix"},
			},
			wantMatcher: StringMatcher{prefixMatch: newStringP("prefix")},
		},
		{
			desc: "happy case prefix ignore case",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "PREFIX"},
				IgnoreCase:   true,
			},
			wantMatcher: StringMatcher{
				prefixMatch: newStringP("prefix"),
				ignoreCase:  true,
			},
		},
		{
			desc: "happy case suffix",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "suffix"},
			},
			wantMatcher: StringMatcher{suffixMatch: newStringP("suffix")},
		},
		{
			desc: "happy case suffix ignore case",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "SUFFIX"},
				IgnoreCase:   true,
			},
			wantMatcher: StringMatcher{
				suffixMatch: newStringP("suffix"),
				ignoreCase:  true,
			},
		},
		{
			desc: "happy case regex",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_SafeRegex{
					SafeRegex: &v3matcherpb.RegexMatcher{Regex: "good?regex?"},
				},
			},
			wantMatcher: StringMatcher{regexMatch: regexp.MustCompile("good?regex?")},
		},
		{
			desc: "happy case contains",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Contains{Contains: "contains"},
			},
			wantMatcher: StringMatcher{containsMatch: newStringP("contains")},
		},
		{
			desc: "happy case contains ignore case",
			inputProto: &v3matcherpb.StringMatcher{
				MatchPattern: &v3matcherpb.StringMatcher_Contains{Contains: "CONTAINS"},
				IgnoreCase:   true,
			},
			wantMatcher: StringMatcher{
				containsMatch: newStringP("contains"),
				ignoreCase:    true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotMatcher, err := StringMatcherFromProto(test.inputProto)
			if (err != nil) != test.wantErr {
				t.Fatalf("StringMatcherFromProto(%+v) returned err: %v, wantErr: %v", test.inputProto, err, test.wantErr)
			}
			if diff := cmp.Diff(gotMatcher, test.wantMatcher, cmp.AllowUnexported(regexp.Regexp{})); diff != "" {
				t.Fatalf("StringMatcherFromProto(%+v) returned unexpected diff (-got, +want):\n%s", test.inputProto, diff)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	var (
		exactMatcher, _           = StringMatcherFromProto(&v3matcherpb.StringMatcher{MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "exact"}})
		prefixMatcher, _          = StringMatcherFromProto(&v3matcherpb.StringMatcher{MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "prefix"}})
		suffixMatcher, _          = StringMatcherFromProto(&v3matcherpb.StringMatcher{MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "suffix"}})
		regexMatcher, _           = StringMatcherFromProto(&v3matcherpb.StringMatcher{MatchPattern: &v3matcherpb.StringMatcher_SafeRegex{SafeRegex: &v3matcherpb.RegexMatcher{Regex: "good?regex?"}}})
		containsMatcher, _        = StringMatcherFromProto(&v3matcherpb.StringMatcher{MatchPattern: &v3matcherpb.StringMatcher_Contains{Contains: "contains"}})
		exactMatcherIgnoreCase, _ = StringMatcherFromProto(&v3matcherpb.StringMatcher{
			MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "exact"},
			IgnoreCase:   true,
		})
		prefixMatcherIgnoreCase, _ = StringMatcherFromProto(&v3matcherpb.StringMatcher{
			MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "prefix"},
			IgnoreCase:   true,
		})
		suffixMatcherIgnoreCase, _ = StringMatcherFromProto(&v3matcherpb.StringMatcher{
			MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "suffix"},
			IgnoreCase:   true,
		})
		containsMatcherIgnoreCase, _ = StringMatcherFromProto(&v3matcherpb.StringMatcher{
			MatchPattern: &v3matcherpb.StringMatcher_Contains{Contains: "contains"},
			IgnoreCase:   true,
		})
	)

	tests := []struct {
		desc      string
		matcher   StringMatcher
		input     string
		wantMatch bool
	}{
		{
			desc:      "exact match success",
			matcher:   exactMatcher,
			input:     "exact",
			wantMatch: true,
		},
		{
			desc:    "exact match failure",
			matcher: exactMatcher,
			input:   "not-exact",
		},
		{
			desc:      "exact match success with ignore case",
			matcher:   exactMatcherIgnoreCase,
			input:     "EXACT",
			wantMatch: true,
		},
		{
			desc:    "exact match failure with ignore case",
			matcher: exactMatcherIgnoreCase,
			input:   "not-exact",
		},
		{
			desc:      "prefix match success",
			matcher:   prefixMatcher,
			input:     "prefixIsHere",
			wantMatch: true,
		},
		{
			desc:    "prefix match failure",
			matcher: prefixMatcher,
			input:   "not-prefix",
		},
		{
			desc:      "prefix match success with ignore case",
			matcher:   prefixMatcherIgnoreCase,
			input:     "PREFIXisHere",
			wantMatch: true,
		},
		{
			desc:    "prefix match failure with ignore case",
			matcher: prefixMatcherIgnoreCase,
			input:   "not-PREFIX",
		},
		{
			desc:      "suffix match success",
			matcher:   suffixMatcher,
			input:     "hereIsThesuffix",
			wantMatch: true,
		},
		{
			desc:    "suffix match failure",
			matcher: suffixMatcher,
			input:   "suffix-is-not-here",
		},
		{
			desc:      "suffix match success with ignore case",
			matcher:   suffixMatcherIgnoreCase,
			input:     "hereIsTheSuFFix",
			wantMatch: true,
		},
		{
			desc:    "suffix match failure with ignore case",
			matcher: suffixMatcherIgnoreCase,
			input:   "SUFFIX-is-not-here",
		},
		{
			desc:      "regex match success",
			matcher:   regexMatcher,
			input:     "goodregex",
			wantMatch: true,
		},
		{
			desc:    "regex match failure",
			matcher: regexMatcher,
			input:   "regex-is-not-here",
		},
		{
			desc:      "contains match success",
			matcher:   containsMatcher,
			input:     "IScontainsHERE",
			wantMatch: true,
		},
		{
			desc:    "contains match failure",
			matcher: containsMatcher,
			input:   "con-tains-is-not-here",
		},
		{
			desc:      "contains match success with ignore case",
			matcher:   containsMatcherIgnoreCase,
			input:     "isCONTAINShere",
			wantMatch: true,
		},
		{
			desc:    "contains match failure with ignore case",
			matcher: containsMatcherIgnoreCase,
			input:   "CON-TAINS-is-not-here",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if gotMatch := test.matcher.Match(test.input); gotMatch != test.wantMatch {
				t.Errorf("StringMatcher.Match(%s) returned %v, want %v", test.input, gotMatch, test.wantMatch)
			}
		})
	}
}

func newStringP(s string) *string {
	return &s
}
