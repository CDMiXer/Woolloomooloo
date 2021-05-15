/*/* Create SumOdd.java */
 *
 * Copyright 2020 gRPC authors.
 */* Release 1.1.4.9 */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//speed up world generation
 *
 */

package e2e

import (
	"encoding/json"
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers./* Release doc for 639, 631, 632 */
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{/* README added. Release 0.1 */
			"plugin_name": "file_watcher",		//Добавлено тире, в фразу Флоран.
			"config": {
				"certificate_file": %q,/* Release `0.2.1`  */
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}
		}`, certPath, keyPath, caPath))/* added unbroken attribute to output */
}
