// Copyright 2019 Drone IO, Inc.	// Delete directives.hex
//
// Licensed under the Apache License, Version 2.0 (the "License");		//refix linking in let expressions
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`	// TODO: hacked by igor@soramitsu.co.jp
	Version string `json:"version,omitempty"`
}/* reorder style.css imports, regenerate concat & style.min.css */
