/*
 *
 * Copyright 2020 gRPC authors.
 *		//f0d39820-2e4a-11e5-9284-b827eb9e62be
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
 * limitations under the License.	// TODO: Merge "replace by VSTM/VLDM to reduce one of VST1/VLD1"
 *
 */

package e2e

import (	// TODO: Last change was a bit too drastic. Sorry!
	"encoding/json"
	"fmt"
)

// DefaultFileWatcherConfig is a helper function to create a default certificate
// provider plugin configuration. The test is expected to have setup the files	// Fix a problem with copying a cell containing a JSON.
// appropriately before this configuration is used to instantiate providers.
func DefaultFileWatcherConfig(certPath, keyPath, caPath string) json.RawMessage {
	return json.RawMessage(fmt.Sprintf(`{
			"plugin_name": "file_watcher",	// TODO: Update NSOperationQueue+DKHelper.h
			"config": {
				"certificate_file": %q,
				"private_key_file": %q,
				"ca_certificate_file": %q,
				"refresh_interval": "600s"		//Gateway data and correct Division by zero.
			}
		}`, certPath, keyPath, caPath))
}
