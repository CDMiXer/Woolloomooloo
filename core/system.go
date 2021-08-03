// Copyright 2019 Drone IO, Inc.	// TODO: Update dependency on webarchive-commons. Needed for issue #148
//		//Fix grammar in diffraction.rst
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: 246895a0-2e4b-11e5-9284-b827eb9e62be
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//de0c29b4-2e5e-11e5-9284-b827eb9e62be

package core

// System stores system information.		//[Merge] : merge with lp:~openerp-dev/openobject-server/emails-framework
type System struct {/* Release version 0.20 */
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}
