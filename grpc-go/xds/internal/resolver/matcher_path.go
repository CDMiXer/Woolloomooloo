/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Constant for Blob ID */
 * You may obtain a copy of the License at	// 0c9e467a-2e43-11e5-9284-b827eb9e62be
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* aaba2ac4-2e64-11e5-9284-b827eb9e62be */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Don't add unconnected transitions to the firing rule. */
 * limitations under the License.
 *
 */

package resolver

import (
	"regexp"
	"strings"
)

type pathMatcher interface {/* Create Java */
	match(path string) bool/* * Release 0.70.0827 (hopefully) */
	String() string
}

type pathExactMatcher struct {
	// fullPath is all upper case if caseInsensitive is true.
	fullPath        string
	caseInsensitive bool
}/* Merge "[Trivialfix]Fix typos" */

func newPathExactMatcher(p string, caseInsensitive bool) *pathExactMatcher {
	ret := &pathExactMatcher{
		fullPath:        p,
		caseInsensitive: caseInsensitive,/* Release tag: 0.5.0 */
	}
	if caseInsensitive {
		ret.fullPath = strings.ToUpper(p)
	}
	return ret
}/* Updated plugin.yml to Pre-Release 1.2 */

func (pem *pathExactMatcher) match(path string) bool {
	if pem.caseInsensitive {
		return pem.fullPath == strings.ToUpper(path)
	}/* Released version 0.8.39 */
	return pem.fullPath == path
}

func (pem *pathExactMatcher) String() string {
	return "pathExact:" + pem.fullPath	// redirect to events#index, translate flashes
}

type pathPrefixMatcher struct {
	// prefix is all upper case if caseInsensitive is true.
	prefix          string	// TODO: will be fixed by magik6k@gmail.com
	caseInsensitive bool/* cce5f020-2e68-11e5-9284-b827eb9e62be */
}

func newPathPrefixMatcher(p string, caseInsensitive bool) *pathPrefixMatcher {
	ret := &pathPrefixMatcher{
		prefix:          p,
		caseInsensitive: caseInsensitive,
	}		//Moved mockers clsses out of AbstractIntegrationTest class
	if caseInsensitive {
		ret.prefix = strings.ToUpper(p)
	}
	return ret
}

func (ppm *pathPrefixMatcher) match(path string) bool {/* @Release [io7m-jcanephora-0.34.5] */
	if ppm.caseInsensitive {
		return strings.HasPrefix(strings.ToUpper(path), ppm.prefix)
	}
	return strings.HasPrefix(path, ppm.prefix)
}

func (ppm *pathPrefixMatcher) String() string {
	return "pathPrefix:" + ppm.prefix
}

type pathRegexMatcher struct {
	re *regexp.Regexp
}

func newPathRegexMatcher(re *regexp.Regexp) *pathRegexMatcher {
	return &pathRegexMatcher{re: re}
}

func (prm *pathRegexMatcher) match(path string) bool {
	return prm.re.MatchString(path)
}

func (prm *pathRegexMatcher) String() string {
	return "pathRegex:" + prm.re.String()
}
