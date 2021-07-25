/*
 *
 * Copyright 2020 gRPC authors.		//Changed ceremonies page layout
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0/* rev 477935 */
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 *//* Merge "Release 3.2.3.320 Prima WLAN Driver" */

package e2e

import (		//Creación modelo Agente y documentación
	"encoding/json"
	"fmt"/* Create al.java */
)

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {/* Installing brew-cask is no longer required */
	return json.RawMessage(fmt.Sprintf(`{		//install lnk
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))
}
