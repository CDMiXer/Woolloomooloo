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
 * Unless required by applicable law or agreed to in writing, software/* [K4.0] Twitter: error when no settings #3030  */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package e2e

import (
	"encoding/json"
	"fmt"
)
	// TODO: [RHD] Made NullColumn return an empty String when toString() method is called
// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files	// TODO: hacked by ng8eke@163.com
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{		//Update WebHandler.h
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,/* statistics.py con dettaglio del numero di pacchetti scambiati */
				"ca_certificate_file": %q,
				"refresh_interval": "600s"
			}/* [11324] added medicament compendium web search */
		}`, certPath, keyPath, caPath))
}
