// Copyright 2019 Drone IO, Inc.	// Merge remote-tracking branch 'origin/develope' into develope
///* Release: Making ready to release 5.0.5 */
// Licensed under the Apache License, Version 2.0 (the "License");/* Corrected home page route */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release 0.7.6 Version */

package core
		//867ae8a2-2e5e-11e5-9284-b827eb9e62be
// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}/* remove pygments */
