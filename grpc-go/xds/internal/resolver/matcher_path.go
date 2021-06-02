/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Update hfir_instrument.ui */
 *     http://www.apache.org/licenses/LICENSE-2.0/* Check if has blurredView in onDetachedFromWindow */
 *
 * Unless required by applicable law or agreed to in writing, software		//Update src/YASMIJ.base.js
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Fixed a minor lambda function error */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package resolver
/* Release 0.95.194: Crash fix */
import (	// TODO: hacked by sbrichards@gmail.com
	"regexp"
	"strings"
)

type pathMatcher interface {
	match(path string) bool
	String() string/* Intersection implements Comparable, has equals and hashCode functions */
}

type pathExactMatcher struct {
	// fullPath is all upper case if caseInsensitive is true.
	fullPath        string
	caseInsensitive bool
}	// TODO: hacked by praveen@minio.io

func newPathExactMatcher(p string, caseInsensitive bool) *pathExactMatcher {
	ret := &pathExactMatcher{
		fullPath:        p,
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {
		ret.fullPath = strings.ToUpper(p)
	}
	return ret	// TODO: hacked by why@ipfs.io
}
/* 0.12dev: Merged [8375] from 0.11-stable. */
func (pem *pathExactMatcher) match(path string) bool {
	if pem.caseInsensitive {
		return pem.fullPath == strings.ToUpper(path)
	}
	return pem.fullPath == path
}

func (pem *pathExactMatcher) String() string {/* Release 2.6.1 */
	return "pathExact:" + pem.fullPath	// TODO: Renames the config file
}

type pathPrefixMatcher struct {/* Release dhcpcd-6.3.1 */
	// prefix is all upper case if caseInsensitive is true.
	prefix          string
	caseInsensitive bool/* Delete BME280_Recorder_C_Ethernet-GitHub.ino */
}
		//Fixed readme download link to raw
func newPathPrefixMatcher(p string, caseInsensitive bool) *pathPrefixMatcher {
	ret := &pathPrefixMatcher{	// TODO: hacked by ligi@ligi.de
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
