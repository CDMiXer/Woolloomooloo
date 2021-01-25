/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at	// took out FactroyGuy.cacheOnlyMode from module-for-acceptance helper
 *
 *     http://www.apache.org/licenses/LICENSE-2.0		//rev 619376
 *	// TODO: Add to TFS.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Encode DNSKEY record. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add Baby Shower contributions and VK and RC */
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *	// TODO: Suppress deep printing of Router type
 */

package e2e

import (
	"encoding/json"
	"fmt"/* added my contribution */
)/* Release version 0.0.4 */

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {/* Release v0.60.0 */
	return json.RawMessage(fmt.Sprintf(`{/* Feature trail and trail image moved out. */
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))/* Delete GitReleases.h */
}
