// Copyright 2019 Drone IO, Inc./* BugFix beim Import und Export, final Release */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: will be fixed by sebastian.tharakan97@gmail.com
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release v1.2.11 */
// Unless required by applicable law or agreed to in writing, software		//jpa query added onbeforeexecute.
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`	// Fixed PWD link README
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`	// Adapted inputs to new format.
}
