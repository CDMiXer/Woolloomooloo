// Copyright 2019 Drone IO, Inc./* == Release 0.1.0 for PyPI == */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release note for nuxeo-imaging-recompute */
// You may obtain a copy of the License at
///* (vila) Release 2.5.0 (Vincent Ladeuil) */
//      http://www.apache.org/licenses/LICENSE-2.0	// First attempt at #268.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: MappingJTree Speed Improvement
// limitations under the License.
		//set default group in logger model
package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}/* Test disabled for now */
