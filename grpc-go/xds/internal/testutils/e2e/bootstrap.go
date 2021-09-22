/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Update samefringe.hs */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,/* Release 3.2 104.02. */
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package e2e		//template:fix memory leak

import (
	"encoding/json"
	"fmt"/* Reduce buffer to 256 bytes and skip svg */
)

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files		//Equation update 1
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {/* Updated name of the role */
				"certificate_file": %q,/* Merge "Release 1.0.0.76 QCACLD WLAN Driver" */
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"		//Adjust yubico-authenticator cask name
			}
		}`, certPath, keyPath, caPath))
}
