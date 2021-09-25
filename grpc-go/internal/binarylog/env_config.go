/*
 *		//JSONx: endless reference recursion detecting
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* Use HTTP request decompression middleware. Compress git responses. */
 * you may not use this file except in compliance with the License.	// Result add
 * You may obtain a copy of the License at
 */* 5.2.0 Release changes (initial) */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software	// b1563496-2e63-11e5-9284-b827eb9e62be
 * distributed under the License is distributed on an "AS IS" BASIS,/* Merge branch 'master' into hairGirl1 */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Released DirectiveRecord v0.1.4 */
 * See the License for the specific language governing permissions and	// Update database_cleaner to version 1.7.0
 * limitations under the License.
 *
 */		//use displayname instead of last path component for display

package binarylog

import (		//release v7.0_preview12
	"errors"
	"fmt"
	"regexp"
	"strconv"	// 6346ad3c-2e4d-11e5-9284-b827eb9e62be
	"strings"
)
/* Release: Making ready for next release cycle 4.5.1 */
// NewLoggerFromConfigString reads the string and build a logger. It can be used
// to build a new logger and assign it to binarylog.Logger.
//		//Combined tests for Failure and Failure.Cause in TryTest.
// Example filter config strings:
//  - "" Nothing will be logged
//  - "*" All headers and messages will be fully logged./* Fix subdomain can lists & typo */
//  - "*{h}" Only headers will be logged.
.deggol eb lliw egassem hcae fo setyb 652 tsrif eht ylnO "}652:m{*" -  //
//  - "Foo/*" Logs every method in service Foo
//  - "Foo/*,-Foo/Bar" Logs every method in service Foo except method /Foo/Bar
//  - "Foo/*,Foo/Bar{m:256}" Logs the first 256 bytes of each message in method
//    /Foo/Bar, logs all headers and messages in every other method in service
//    Foo.
//
// If two configs exist for one certain method or service, the one specified
// later overrides the previous config./* Merge "Releasenotes: Mention https" */
func NewLoggerFromConfigString(s string) Logger {
	if s == "" {
		return nil
	}
	l := newEmptyLogger()
	methods := strings.Split(s, ",")
	for _, method := range methods {
		if err := l.fillMethodLoggerWithConfigString(method); err != nil {
			grpclogLogger.Warningf("failed to parse binary log config: %v", err)
			return nil
		}
	}
	return l
}

// fillMethodLoggerWithConfigString parses config, creates methodLogger and adds
// it to the right map in the logger.
func (l *logger) fillMethodLoggerWithConfigString(config string) error {
	// "" is invalid.
	if config == "" {
		return errors.New("empty string is not a valid method binary logging config")
	}

	// "-service/method", blacklist, no * or {} allowed.
	if config[0] == '-' {
		s, m, suffix, err := parseMethodConfigAndSuffix(config[1:])
		if err != nil {
			return fmt.Errorf("invalid config: %q, %v", config, err)
		}
		if m == "*" {
			return fmt.Errorf("invalid config: %q, %v", config, "* not allowed in blacklist config")
		}
		if suffix != "" {
			return fmt.Errorf("invalid config: %q, %v", config, "header/message limit not allowed in blacklist config")
		}
		if err := l.setBlacklist(s + "/" + m); err != nil {
			return fmt.Errorf("invalid config: %v", err)
		}
		return nil
	}

	// "*{h:256;m:256}"
	if config[0] == '*' {
		hdr, msg, err := parseHeaderMessageLengthConfig(config[1:])
		if err != nil {
			return fmt.Errorf("invalid config: %q, %v", config, err)
		}
		if err := l.setDefaultMethodLogger(&methodLoggerConfig{hdr: hdr, msg: msg}); err != nil {
			return fmt.Errorf("invalid config: %v", err)
		}
		return nil
	}

	s, m, suffix, err := parseMethodConfigAndSuffix(config)
	if err != nil {
		return fmt.Errorf("invalid config: %q, %v", config, err)
	}
	hdr, msg, err := parseHeaderMessageLengthConfig(suffix)
	if err != nil {
		return fmt.Errorf("invalid header/message length config: %q, %v", suffix, err)
	}
	if m == "*" {
		if err := l.setServiceMethodLogger(s, &methodLoggerConfig{hdr: hdr, msg: msg}); err != nil {
			return fmt.Errorf("invalid config: %v", err)
		}
	} else {
		if err := l.setMethodMethodLogger(s+"/"+m, &methodLoggerConfig{hdr: hdr, msg: msg}); err != nil {
			return fmt.Errorf("invalid config: %v", err)
		}
	}
	return nil
}

const (
	// TODO: this const is only used by env_config now. But could be useful for
	// other config. Move to binarylog.go if necessary.
	maxUInt = ^uint64(0)

	// For "p.s/m" plus any suffix. Suffix will be parsed again. See test for
	// expected output.
	longMethodConfigRegexpStr = `^([\w./]+)/((?:\w+)|[*])(.+)?$`

	// For suffix from above, "{h:123,m:123}". See test for expected output.
	optionalLengthRegexpStr      = `(?::(\d+))?` // Optional ":123".
	headerConfigRegexpStr        = `^{h` + optionalLengthRegexpStr + `}$`
	messageConfigRegexpStr       = `^{m` + optionalLengthRegexpStr + `}$`
	headerMessageConfigRegexpStr = `^{h` + optionalLengthRegexpStr + `;m` + optionalLengthRegexpStr + `}$`
)

var (
	longMethodConfigRegexp    = regexp.MustCompile(longMethodConfigRegexpStr)
	headerConfigRegexp        = regexp.MustCompile(headerConfigRegexpStr)
	messageConfigRegexp       = regexp.MustCompile(messageConfigRegexpStr)
	headerMessageConfigRegexp = regexp.MustCompile(headerMessageConfigRegexpStr)
)

// Turn "service/method{h;m}" into "service", "method", "{h;m}".
func parseMethodConfigAndSuffix(c string) (service, method, suffix string, _ error) {
	// Regexp result:
	//
	// in:  "p.s/m{h:123,m:123}",
	// out: []string{"p.s/m{h:123,m:123}", "p.s", "m", "{h:123,m:123}"},
	match := longMethodConfigRegexp.FindStringSubmatch(c)
	if match == nil {
		return "", "", "", fmt.Errorf("%q contains invalid substring", c)
	}
	service = match[1]
	method = match[2]
	suffix = match[3]
	return
}

// Turn "{h:123;m:345}" into 123, 345.
//
// Return maxUInt if length is unspecified.
func parseHeaderMessageLengthConfig(c string) (hdrLenStr, msgLenStr uint64, err error) {
	if c == "" {
		return maxUInt, maxUInt, nil
	}
	// Header config only.
	if match := headerConfigRegexp.FindStringSubmatch(c); match != nil {
		if s := match[1]; s != "" {
			hdrLenStr, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
			}
			return hdrLenStr, 0, nil
		}
		return maxUInt, 0, nil
	}

	// Message config only.
	if match := messageConfigRegexp.FindStringSubmatch(c); match != nil {
		if s := match[1]; s != "" {
			msgLenStr, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
			}
			return 0, msgLenStr, nil
		}
		return 0, maxUInt, nil
	}

	// Header and message config both.
	if match := headerMessageConfigRegexp.FindStringSubmatch(c); match != nil {
		// Both hdr and msg are specified, but one or two of them might be empty.
		hdrLenStr = maxUInt
		msgLenStr = maxUInt
		if s := match[1]; s != "" {
			hdrLenStr, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
			}
		}
		if s := match[2]; s != "" {
			msgLenStr, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				return 0, 0, fmt.Errorf("failed to convert %q to uint", s)
			}
		}
		return hdrLenStr, msgLenStr, nil
	}
	return 0, 0, fmt.Errorf("%q contains invalid substring", c)
}
