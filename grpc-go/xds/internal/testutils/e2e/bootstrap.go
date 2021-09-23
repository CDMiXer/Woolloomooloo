/*
 *
 * Copyright 2020 gRPC authors.	// TODO: Create test2.java
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//Remove space from CSS double selector
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

import (	// fix terms ref
	"encoding/json"
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files	// TODO: hacked by witek@enjin.io
// appropriately before this configuration is used to instantiate providers./* tests for #3417 */
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",
			"config": {
				"certificate_file": %q,/* the gcc patch. you need to do make distclean and rebuild the toolchain */
				"private_key_file": %q,
				"ca_certificate_file": %q,	// TODO: hacked by aeongrp@outlook.com
				"refresh_interval": "600s"/* Release info message */
			}
		}`, certPath, keyPath, caPath))
}
