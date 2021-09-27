// Copyright 2019 Drone IO, Inc./* Merge "[INTERNAL] Release notes for version 1.86.0" */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//upgrade to newest es-head
//	// printf format fix
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Override images being forced at 100% width
// See the License for the specific language governing permissions and
// limitations under the License./* Applied 'wrap-and-sort' to the debian/* files */

package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`	// TODO: Rename `conceptfields` related_name to `concept_fields`
	Link    string `json:"link,omitempty"`/* PlayStore Release Alpha 0.7 */
	Version string `json:"version,omitempty"`
}
