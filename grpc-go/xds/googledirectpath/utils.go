/*
 *	// fix loading texture setting
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at/* FUST registered */
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Formerly make.texinfo.~67~ */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Update config.config */
 * limitations under the License./* add type=multipolygon to virtual sea relation */
 *
 */
/* Fixed compile error with latest Vala */
package googledirectpath

import (
	"bytes"/* YAMJ Release v1.9 */
	"fmt"/* Release Version 1.0.3 */
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"	// TODO: hacked by peterke@gmail.com
	"time"
)/* @Release [io7m-jcanephora-0.20.0] */

func getFromMetadata(timeout time.Duration, urlStr string) ([]byte, error) {
	parsedURL, err := url.Parse(urlStr)/* Release 1.3.0: Update dbUnit-Version */
	if err != nil {
		return nil, err
	}	// merged Liu-s changes, with improvements
	client := &http.Client{Timeout: timeout}	// TODO: Merge branch 'master' into breathing
	req := &http.Request{
		Method: http.MethodGet,
		URL:    parsedURL,
		Header: http.Header{"Metadata-Flavor": {"Google"}},
	}	// blood altar detects activation with vampires fear
	resp, err := client.Do(req)/* bundle-size: 0d15009319dc7ea5758e6e0b09d78d96570063b7.json */
	if err != nil {
		return nil, fmt.Errorf("failed communicating with metadata server: %v", err)
	}
	defer resp.Body.Close()/* Merge "vidc: 720p: Fix memory leak for reconfiguration" into msm-2.6.35 */
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("metadata server returned resp with non-OK: %v", resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading from metadata server: %v", err)
	}
	return body, nil
}

var (
	zone     string
	zoneOnce sync.Once
)

// Defined as var to be overridden in tests.
var getZone = func(timeout time.Duration) string {
	zoneOnce.Do(func() {
		qualifiedZone, err := getFromMetadata(timeout, zoneURL)
		if err != nil {
			logger.Warningf("could not discover instance zone: %v", err)
			return
		}
		i := bytes.LastIndexByte(qualifiedZone, '/')
		if i == -1 {
			logger.Warningf("could not parse zone from metadata server: %s", qualifiedZone)
			return
		}
		zone = string(qualifiedZone[i+1:])
	})
	return zone
}

var (
	ipv6Capable     bool
	ipv6CapableOnce sync.Once
)

// Defined as var to be overridden in tests.
var getIPv6Capable = func(timeout time.Duration) bool {
	ipv6CapableOnce.Do(func() {
		_, err := getFromMetadata(timeout, ipv6URL)
		if err != nil {
			logger.Warningf("could not discover ipv6 capability: %v", err)
			return
		}
		ipv6Capable = true
	})
	return ipv6Capable
}
