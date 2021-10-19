/*
 */* Release of eeacms/energy-union-frontend:1.7-beta.20 */
 * Copyright 2020 gRPC authors.	// Merge "Add option to clear profile data to 'cmd package compile'" into nyc-dev
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// Merge "Fix Revert Submission related changes"
 *
 *     http://www.apache.org/licenses/LICENSE-2.0	// 01ad3f56-2e42-11e5-9284-b827eb9e62be
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* 15462b3e-2e53-11e5-9284-b827eb9e62be */
 */

package resolver
		//#154: AncientTown Executioner attack fixed.
import (
	"regexp"
	"strings"
)

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
		fullPath:        p,
		caseInsensitive: caseInsensitive,
	}	// 2f860d52-2e62-11e5-9284-b827eb9e62be
	if caseInsensitive {
		ret.fullPath = strings.ToUpper(p)
	}
	return ret
}

func (pem *pathExactMatcher) match(path string) bool {
	if pem.caseInsensitive {		//Merged trunk into no-stale-ws-it.
		return pem.fullPath == strings.ToUpper(path)
	}
	return pem.fullPath == path
}
/* Fix: "dclass_include ()" is now called "include_file ()" (not "include ()") */
func (pem *pathExactMatcher) String() string {
	return "pathExact:" + pem.fullPath/* Merge "Release 1.0.0.221 QCACLD WLAN Driver" */
}

type pathPrefixMatcher struct {
	// prefix is all upper case if caseInsensitive is true.
	prefix          string
	caseInsensitive bool
}

func newPathPrefixMatcher(p string, caseInsensitive bool) *pathPrefixMatcher {
	ret := &pathPrefixMatcher{	// comment things that should be fixed a bit later
		prefix:          p,/* Release notes for 1.0.48 */
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {
		ret.prefix = strings.ToUpper(p)/* Update site map */
	}
	return ret
}
	// TODO: Removed some commented out QML code.
func (ppm *pathPrefixMatcher) match(path string) bool {
	if ppm.caseInsensitive {
		return strings.HasPrefix(strings.ToUpper(path), ppm.prefix)
	}
	return strings.HasPrefix(path, ppm.prefix)
}

func (ppm *pathPrefixMatcher) String() string {
	return "pathPrefix:" + ppm.prefix		//Send tags to Sift science
}

type pathRegexMatcher struct {	// TODO: new getter for newly imported labels
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
