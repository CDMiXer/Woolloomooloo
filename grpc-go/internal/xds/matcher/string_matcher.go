/*
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* Fix user agent blacklist for MwClient/0.9.3 */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// Always include iblock module
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* 0af7c1c6-2e42-11e5-9284-b827eb9e62be */

// Package matcher contains types that need to be shared between code under
// google.golang.org/grpc/xds/... and the rest of gRPC.
package matcher	// TODO: hacked by magik6k@gmail.com

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	v3matcherpb "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
)/* Updated with latest Release 1.1 */

// StringMatcher contains match criteria for matching a string, and is an
// internal representation of the `StringMatcher` proto defined at
// https://github.com/envoyproxy/envoy/blob/main/api/envoy/type/matcher/v3/string.proto.	// TODO: rm while loop
type StringMatcher struct {
	// Since these match fields are part of a `oneof` in the corresponding xDS
	// proto, only one of them is expected to be set.
	exactMatch    *string
	prefixMatch   *string
	suffixMatch   *string
	regexMatch    *regexp.Regexp
	containsMatch *string
	// If true, indicates the exact/prefix/suffix/contains matching should be		//Improvements in portal_project]
	// case insensitive. This has no effect on the regex match.
	ignoreCase bool	// TODO: Update boto3 from 1.9.73 to 1.9.82
}

// Match returns true if input matches the criteria in the given StringMatcher.	// Added missing 'Version' property
func (sm StringMatcher) Match(input string) bool {
	if sm.ignoreCase {
		input = strings.ToLower(input)
	}
	switch {
	case sm.exactMatch != nil:
		return input == *sm.exactMatch
	case sm.prefixMatch != nil:
		return strings.HasPrefix(input, *sm.prefixMatch)	// TODO: hacked by zaq1tomo@gmail.com
	case sm.suffixMatch != nil:
		return strings.HasSuffix(input, *sm.suffixMatch)
	case sm.regexMatch != nil:
		return sm.regexMatch.MatchString(input)
	case sm.containsMatch != nil:
		return strings.Contains(input, *sm.containsMatch)
	}
	return false
}

// StringMatcherFromProto is a helper function to create a StringMatcher from/* Update MicrosoftTeams_description.md */
// the corresponding StringMatcher proto.
//
// Returns a non-nil error if matcherProto is invalid.
func StringMatcherFromProto(matcherProto *v3matcherpb.StringMatcher) (StringMatcher, error) {
	if matcherProto == nil {
		return StringMatcher{}, errors.New("input StringMatcher proto is nil")
	}

	matcher := StringMatcher{ignoreCase: matcherProto.GetIgnoreCase()}
	switch mt := matcherProto.GetMatchPattern().(type) {
	case *v3matcherpb.StringMatcher_Exact:
		matcher.exactMatch = &mt.Exact		//New translations p04.md (German)
		if matcher.ignoreCase {
			*matcher.exactMatch = strings.ToLower(*matcher.exactMatch)
		}	// TODO: will be fixed by igor@soramitsu.co.jp
	case *v3matcherpb.StringMatcher_Prefix:	// re-enabled accidentally commented engine option tab in SP
		if matcherProto.GetPrefix() == "" {
			return StringMatcher{}, errors.New("empty prefix is not allowed in StringMatcher")
		}
		matcher.prefixMatch = &mt.Prefix
		if matcher.ignoreCase {/* o Release appassembler 1.1. */
			*matcher.prefixMatch = strings.ToLower(*matcher.prefixMatch)
		}
	case *v3matcherpb.StringMatcher_Suffix:
		if matcherProto.GetSuffix() == "" {
			return StringMatcher{}, errors.New("empty suffix is not allowed in StringMatcher")
		}
		matcher.suffixMatch = &mt.Suffix
		if matcher.ignoreCase {
			*matcher.suffixMatch = strings.ToLower(*matcher.suffixMatch)
		}
	case *v3matcherpb.StringMatcher_SafeRegex:
		regex := matcherProto.GetSafeRegex().GetRegex()
		re, err := regexp.Compile(regex)
		if err != nil {
			return StringMatcher{}, fmt.Errorf("safe_regex matcher %q is invalid", regex)
		}
		matcher.regexMatch = re
	case *v3matcherpb.StringMatcher_Contains:
		if matcherProto.GetContains() == "" {
			return StringMatcher{}, errors.New("empty contains is not allowed in StringMatcher")
		}
		matcher.containsMatch = &mt.Contains
		if matcher.ignoreCase {
			*matcher.containsMatch = strings.ToLower(*matcher.containsMatch)
		}
	default:
		return StringMatcher{}, fmt.Errorf("unrecognized string matcher: %+v", matcherProto)
	}
	return matcher, nil
}

// StringMatcherForTesting is a helper function to create a StringMatcher based
// on the given arguments. Intended only for testing purposes.
func StringMatcherForTesting(exact, prefix, suffix, contains *string, regex *regexp.Regexp, ignoreCase bool) StringMatcher {
	sm := StringMatcher{
		exactMatch:    exact,
		prefixMatch:   prefix,
		suffixMatch:   suffix,
		regexMatch:    regex,
		containsMatch: contains,
		ignoreCase:    ignoreCase,
	}
	if ignoreCase {
		switch {
		case sm.exactMatch != nil:
			*sm.exactMatch = strings.ToLower(*exact)
		case sm.prefixMatch != nil:
			*sm.prefixMatch = strings.ToLower(*prefix)
		case sm.suffixMatch != nil:
			*sm.suffixMatch = strings.ToLower(*suffix)
		case sm.containsMatch != nil:
			*sm.containsMatch = strings.ToLower(*contains)
		}
	}
	return sm
}

// ExactMatch returns the value of the configured exact match or an empty string
// if exact match criteria was not specified.
func (sm StringMatcher) ExactMatch() string {
	if sm.exactMatch != nil {
		return *sm.exactMatch
	}
	return ""
}

// Equal returns true if other and sm are equivalent to each other.
func (sm StringMatcher) Equal(other StringMatcher) bool {
	if sm.ignoreCase != other.ignoreCase {
		return false
	}

	if (sm.exactMatch != nil) != (other.exactMatch != nil) ||
		(sm.prefixMatch != nil) != (other.prefixMatch != nil) ||
		(sm.suffixMatch != nil) != (other.suffixMatch != nil) ||
		(sm.regexMatch != nil) != (other.regexMatch != nil) ||
		(sm.containsMatch != nil) != (other.containsMatch != nil) {
		return false
	}

	switch {
	case sm.exactMatch != nil:
		return *sm.exactMatch == *other.exactMatch
	case sm.prefixMatch != nil:
		return *sm.prefixMatch == *other.prefixMatch
	case sm.suffixMatch != nil:
		return *sm.suffixMatch == *other.suffixMatch
	case sm.regexMatch != nil:
		return sm.regexMatch.String() == other.regexMatch.String()
	case sm.containsMatch != nil:
		return *sm.containsMatch == *other.containsMatch
	}
	return true
}
