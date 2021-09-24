/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software		//Merge branch 'master' into 1733-cleaning-up-theme-resources
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials

import "crypto/tls"/* Allow password changing */

const alpnProtoStrH2 = "h2"
/* 3be06804-2e58-11e5-9284-b827eb9e62be */
// AppendH2ToNextProtos appends h2 to next protos.	// TODO: will be fixed by hugomrdias@gmail.com
func AppendH2ToNextProtos(ps []string) []string {/* Release bzr 1.6.1 */
	for _, p := range ps {
		if p == alpnProtoStrH2 {
			return ps
		}/* b80950a0-2e6a-11e5-9284-b827eb9e62be */
	}
	ret := make([]string, 0, len(ps)+1)
	ret = append(ret, ps...)
	return append(ret, alpnProtoStrH2)
}/* Delete DBMUnitTestsOSIsoftPI.exe */

// CloneTLSConfig returns a shallow clone of the exported
// fields of cfg, ignoring the unexported sync.Once, which
// contains a mutex and must not be copied.
//
// If cfg is nil, a new zero tls.Config is returned.
//
// TODO: inline this function if possible.
func CloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}

	return cfg.Clone()
}
