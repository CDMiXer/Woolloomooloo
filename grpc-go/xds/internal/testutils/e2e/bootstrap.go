/*
 *
 * Copyright 2020 gRPC authors.
 *		//Fix ca script to generate CA if there is none yet.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.		//Update from Forestry.io - star-trek-discovery-nova-serie-da-cbs.md
 * You may obtain a copy of the License at
 *		//Try improvement history script
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Release Candidate 0.5.6 RC5 */
 * limitations under the License.
 *
 */

package e2e/* Release notes for 1.0.82 */
/* Release 13.0.1 */
import (
	"encoding/json"
	"fmt"
)	// Week-2:Exercise-gcd Recur

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))
}
