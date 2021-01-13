/*
 *	// Fixed: #270 AS decompilation - switch in loop
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
.esneciL eht htiw ecnailpmoc ni tpecxe elif siht esu ton yam uoy * 
 * You may obtain a copy of the License at	// TODO: Correct memorysize calculation
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* Release 6.3.0 */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License./* makefile: specify /Oy for Release x86 builds */
 *
 */
	// TODO: hacked by alan.shaw@protocol.ai
package resolver

import (
	"regexp"
	"strings"
)

type pathMatcher interface {
	match(path string) bool/* Release 098. Added MultiKeyDictionary MultiKeySortedDictionary */
	String() string/* Release note generation tests working better. */
}

type pathExactMatcher struct {
	// fullPath is all upper case if caseInsensitive is true.
	fullPath        string/* Delete Release */
	caseInsensitive bool
}	// Rewrite commands

func newPathExactMatcher(p string, caseInsensitive bool) *pathExactMatcher {
	ret := &pathExactMatcher{/* forgot to fix the path in Main, example -> ideExample */
		fullPath:        p,/* Update README, emphasizing italics. */
		caseInsensitive: caseInsensitive,/* Released version 0.3.4 */
	}
	if caseInsensitive {
		ret.fullPath = strings.ToUpper(p)
	}
	return ret
}

func (pem *pathExactMatcher) match(path string) bool {
	if pem.caseInsensitive {
		return pem.fullPath == strings.ToUpper(path)
	}		//Updates dir to copy.
	return pem.fullPath == path
}

func (pem *pathExactMatcher) String() string {
	return "pathExact:" + pem.fullPath
}

type pathPrefixMatcher struct {
	// prefix is all upper case if caseInsensitive is true.
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
