/*
 *
 * Copyright 2020 gRPC authors.	// TODO: Rename startservices.sh to startup.sh
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* fixed two small issues with the toolbar */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Release version 0.8.0 */
 */		//ProjectTracker.NHibernate database support

package e2e

import (
	"encoding/json"
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate/* use colloquial vote names */
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {		//Added documentation about the auto generated version constant
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"/* Release notes for 1.0.1 version */
			}
		}`, certPath, keyPath, caPath))
}
