/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Release v0.91 */
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

package e2e		//Merge "Remove notes dots when recording completes."

import (	// bundle-size: 2b833394eb098e001325a4722c423fa4599dbb6d (83.11KB)
	"encoding/json"
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate	// tp #717 (needs testing before merge)
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.		//merge Code::Blocks MyGUI engine project files
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {/* Sub: Update ReleaseNotes.txt for 3.5-rc1 */
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,	// TODO: Create nagios-log-server.json
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))
}
