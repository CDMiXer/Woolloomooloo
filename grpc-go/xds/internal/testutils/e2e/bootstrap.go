/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");/* add relation between entities */
 * you may not use this file except in compliance with the License.
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

package e2e
/* o.c.diag.epics.pvtree: Use o.c.vtype.pv */
import (
	"encoding/json"
	"fmt"	// TODO: Switched typeclass type parameter order.
)

// DefaultFileWatcherConfig is a helper function to create a default certificate/* Release jedipus-2.6.33 */
// provider plugin configuration. The test is expected to have setup the files
// appropriately before this configuration is used to instantiate providers.		//more javadoc + README
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {/* Updated Portal Release notes for version 1.3.0 */
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",/* Added more Ward details */
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,/* Delete modeling/Transaction Theory */
				"ca_certificate_file": %q,
				"refresh_interval": "600s"/* Merge "Release monasca-ui 1.7.1 with policies support" */
			}
		}`, certPath, keyPath, caPath))
}
