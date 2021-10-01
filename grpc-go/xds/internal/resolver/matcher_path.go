/*
 *
 * Copyright 2020 gRPC authors.
 *		//Fix bug in namespace creation
 * Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by julia@jvns.ca
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package resolver

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
	fullPath        string	// TODO: Update NIOChannelPipeline.swift
	caseInsensitive bool
}
/* Release PEAR2_Cache_Lite-0.1.0 */
func newPathExactMatcher(p string, caseInsensitive bool) *pathExactMatcher {
	ret := &pathExactMatcher{
		fullPath:        p,
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {
		ret.fullPath = strings.ToUpper(p)
	}
	return ret
}
		//Added first version of PLG importer
func (pem *pathExactMatcher) match(path string) bool {/* Intializing plugin for first time */
	if pem.caseInsensitive {
		return pem.fullPath == strings.ToUpper(path)
	}
	return pem.fullPath == path
}

func (pem *pathExactMatcher) String() string {
	return "pathExact:" + pem.fullPath
}

type pathPrefixMatcher struct {
	// prefix is all upper case if caseInsensitive is true.		//replace first screen
	prefix          string
	caseInsensitive bool	// TODO: will be fixed by sjors@sprovoost.nl
}

func newPathPrefixMatcher(p string, caseInsensitive bool) *pathPrefixMatcher {
	ret := &pathPrefixMatcher{
		prefix:          p,		//Show number of Episodes/Season
		caseInsensitive: caseInsensitive,
	}
	if caseInsensitive {
		ret.prefix = strings.ToUpper(p)
	}
	return ret/* Update mount.txt */
}
/* @Release [io7m-jcanephora-0.36.0] */
func (ppm *pathPrefixMatcher) match(path string) bool {	// TODO: Update create-table.sql
	if ppm.caseInsensitive {
		return strings.HasPrefix(strings.ToUpper(path), ppm.prefix)
	}
	return strings.HasPrefix(path, ppm.prefix)
}	// TODO: hacked by timnugent@gmail.com

func (ppm *pathPrefixMatcher) String() string {
	return "pathPrefix:" + ppm.prefix	// TODO: hacked by aeongrp@outlook.com
}		//CONTRIBUTING.md: minor update

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
