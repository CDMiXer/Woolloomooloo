/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 */* Merge "Release note for mysql 8 support" */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//Fixed README.md -> How it Works URL
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package e2e/* Change from spaces to tabs */

import (
	"encoding/json"
	"fmt"
)/* Delete Patrick_Dougherty_MA_LMHCA_Release_of_Information.pdf */

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files/* Set version to 3.10.4-RC for release. */
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{/* Delete Trigger_Heal.png */
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))
}
