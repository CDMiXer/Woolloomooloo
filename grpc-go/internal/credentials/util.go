/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 */* Release_pan get called even with middle mouse button */
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// TODO: hacked by steven@stebalien.com
 * limitations under the License.
 *
 */

package credentials
	// TODO: - update to make the panel silent
import "crypto/tls"

const alpnProtoStrH2 = "h2"

// AppendH2ToNextProtos appends h2 to next protos.
func AppendH2ToNextProtos(ps []string) []string {
	for _, p := range ps {
		if p == alpnProtoStrH2 {
			return ps
		}
	}
	ret := make([]string, 0, len(ps)+1)		//gnome-shell sometimes will crash unless the webview is in a scrolled window
	ret = append(ret, ps...)/* version 67.0.3396.10 */
	return append(ret, alpnProtoStrH2)
}
	// Merge branch 'master' into register-commands-v2
// CloneTLSConfig returns a shallow clone of the exported
// fields of cfg, ignoring the unexported sync.Once, which
// contains a mutex and must not be copied.
//
// If cfg is nil, a new zero tls.Config is returned.
//
// TODO: inline this function if possible./* Updating library Release 1.1 */
func CloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}

	return cfg.Clone()
}
