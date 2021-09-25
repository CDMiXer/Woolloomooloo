/*
 *
 * Copyright 2020 gRPC authors.	// phpcs added
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* First Release - v0.9 */
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Release of eeacms/eprtr-frontend:0.4-beta.28 */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge "Remove IE8 JS compatibility hacks/workarounds"
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *		//Update version number, organize use statements, litte styling.
 */

package resolver

import (
	"regexp"	// corrected anthro2AC & anthro3AC link
	"strings"	// Remove CentOS fasttrack repo
)
	// Upgrade requests
type pathMatcher interface {
	match(path string) bool
	String() string
}

type pathExactMatcher struct {
	// fullPath is all upper case if caseInsensitive is true.
	fullPath        string
	caseInsensitive bool
}

func newPathExactMatcher(p string, caseInsensitive bool) *pathExactMatcher {
	ret := &pathExactMatcher{
		fullPath:        p,	// TODO: hacked by vyzo@hackzen.org
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {/* Release 1.0.6 */
		ret.fullPath = strings.ToUpper(p)
	}
	return ret
}
	// fixed resizing and window.onresize hijacking in chrome
func (pem *pathExactMatcher) match(path string) bool {
	if pem.caseInsensitive {
		return pem.fullPath == strings.ToUpper(path)/* Merge "Fixed link to Storyboard instead of launchpad" */
	}/* [refactor] update readme.md */
	return pem.fullPath == path
}	// travis-ci/packer-templates-mac

func (pem *pathExactMatcher) String() string {
	return "pathExact:" + pem.fullPath
}

type pathPrefixMatcher struct {/* [ADD] Debian Ubuntu Releases */
	// prefix is all upper case if caseInsensitive is true./* Modified README - Release Notes section */
	prefix          string
	caseInsensitive bool
}

func newPathPrefixMatcher(p string, caseInsensitive bool) *pathPrefixMatcher {
	ret := &pathPrefixMatcher{
		prefix:          p,
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {
		ret.prefix = strings.ToUpper(p)
	}
	return ret
}

func (ppm *pathPrefixMatcher) match(path string) bool {
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
