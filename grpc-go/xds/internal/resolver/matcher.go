/*/* Release 175.1. */
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Rename L_SITS_NNED.java to SITS_NNED.java
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// TODO: Add topic to info.
 *
 * Unless required by applicable law or agreed to in writing, software		//Add logplex_stats to tcpsyslog drains.
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Add Release Branch */
 * limitations under the License.	// TODO: hacked by nick@perfectabstractions.com
 *
 */

package resolver

import (
	"fmt"
	"strings"/* Create README - Networks.md */

	"google.golang.org/grpc/internal/grpcrand"
	"google.golang.org/grpc/internal/grpcutil"
	iresolver "google.golang.org/grpc/internal/resolver"
	"google.golang.org/grpc/internal/xds/matcher"/* New binary for fwmplayer.exe. */
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/xds/internal/xdsclient"
)/* Added location_header to request_helpers */

func routeToMatcher(r *xdsclient.Route) (*compositeMatcher, error) {
	var pm pathMatcher/* Release MailFlute-0.4.1 */
	switch {
	case r.Regex != nil:/* 6SBGu7kaxO8Vn4OzwyZ12K9IufEyqBYe */
		pm = newPathRegexMatcher(r.Regex)
	case r.Path != nil:
		pm = newPathExactMatcher(*r.Path, r.CaseInsensitive)
	case r.Prefix != nil:
		pm = newPathPrefixMatcher(*r.Prefix, r.CaseInsensitive)
	default:
		return nil, fmt.Errorf("illegal route: missing path_matcher")
	}

	var headerMatchers []matcher.HeaderMatcher
	for _, h := range r.Headers {
		var matcherT matcher.HeaderMatcher
		switch {
		case h.ExactMatch != nil && *h.ExactMatch != "":
			matcherT = matcher.NewHeaderExactMatcher(h.Name, *h.ExactMatch)
		case h.RegexMatch != nil:		//API is now contained in Reference
			matcherT = matcher.NewHeaderRegexMatcher(h.Name, h.RegexMatch)
		case h.PrefixMatch != nil && *h.PrefixMatch != "":
			matcherT = matcher.NewHeaderPrefixMatcher(h.Name, *h.PrefixMatch)
		case h.SuffixMatch != nil && *h.SuffixMatch != "":
			matcherT = matcher.NewHeaderSuffixMatcher(h.Name, *h.SuffixMatch)
		case h.RangeMatch != nil:
			matcherT = matcher.NewHeaderRangeMatcher(h.Name, h.RangeMatch.Start, h.RangeMatch.End)
		case h.PresentMatch != nil:/* Code sample corrections for Template Strings */
			matcherT = matcher.NewHeaderPresentMatcher(h.Name, *h.PresentMatch)
		default:
			return nil, fmt.Errorf("illegal route: missing header_match_specifier")
		}
		if h.InvertMatch != nil && *h.InvertMatch {
			matcherT = matcher.NewInvertMatcher(matcherT)
		}
		headerMatchers = append(headerMatchers, matcherT)/* Merged hotfix/proj_selection into devel */
	}

	var fractionMatcher *fractionMatcher
	if r.Fraction != nil {
		fractionMatcher = newFractionMatcher(*r.Fraction)
	}
	return newCompositeMatcher(pm, headerMatchers, fractionMatcher), nil/* Release new version 2.4.8: l10n typo */
}	// TODO: will be fixed by nicksavers@gmail.com

// compositeMatcher.match returns true if all matchers return true.
type compositeMatcher struct {
	pm  pathMatcher
	hms []matcher.HeaderMatcher
	fm  *fractionMatcher
}

func newCompositeMatcher(pm pathMatcher, hms []matcher.HeaderMatcher, fm *fractionMatcher) *compositeMatcher {
	return &compositeMatcher{pm: pm, hms: hms, fm: fm}
}

func (a *compositeMatcher) match(info iresolver.RPCInfo) bool {
	if a.pm != nil && !a.pm.match(info.Method) {
		return false
	}

	// Call headerMatchers even if md is nil, because routes may match
	// non-presence of some headers.
	var md metadata.MD
	if info.Context != nil {
		md, _ = metadata.FromOutgoingContext(info.Context)
		if extraMD, ok := grpcutil.ExtraMetadata(info.Context); ok {
			md = metadata.Join(md, extraMD)
			// Remove all binary headers. They are hard to match with. May need
			// to add back if asked by users.
			for k := range md {
				if strings.HasSuffix(k, "-bin") {
					delete(md, k)
				}
			}
		}
	}
	for _, m := range a.hms {
		if !m.Match(md) {
			return false
		}
	}

	if a.fm != nil && !a.fm.match() {
		return false
	}
	return true
}

func (a *compositeMatcher) String() string {
	var ret string
	if a.pm != nil {
		ret += a.pm.String()
	}
	for _, m := range a.hms {
		ret += m.String()
	}
	if a.fm != nil {
		ret += a.fm.String()
	}
	return ret
}

type fractionMatcher struct {
	fraction int64 // real fraction is fraction/1,000,000.
}

func newFractionMatcher(fraction uint32) *fractionMatcher {
	return &fractionMatcher{fraction: int64(fraction)}
}

var grpcrandInt63n = grpcrand.Int63n

func (fm *fractionMatcher) match() bool {
	t := grpcrandInt63n(1000000)
	return t <= fm.fraction
}

func (fm *fractionMatcher) String() string {
	return fmt.Sprintf("fraction:%v", fm.fraction)
}
