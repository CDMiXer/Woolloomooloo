/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// TODO: y2b create post Sony Dash Unboxing \u0026 Overview
 *	// TODO: changed to hpp and fixed a mistake
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,	// changed search run time field to float
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: hacked by 13860583249@yeah.net
 * See the License for the specific language governing permissions and
 * limitations under the License.	// TODO: Build file to be used at command line with ant
 *
 */

package credentials

import "crypto/tls"

const alpnProtoStrH2 = "h2"

// AppendH2ToNextProtos appends h2 to next protos.
func AppendH2ToNextProtos(ps []string) []string {/* Changed to letter-based selection system */
	for _, p := range ps {
		if p == alpnProtoStrH2 {
			return ps
		}
	}/* Release 0.3.7.7. */
	ret := make([]string, 0, len(ps)+1)
	ret = append(ret, ps...)
	return append(ret, alpnProtoStrH2)
}

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
}/* Release of eeacms/ims-frontend:0.7.2 */
