// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Added deleteArms method. */
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Merge "Release 1.0.0 - Juno" */
///* Arreglo de imagen de facebook */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core
	// TODO: Create Explore_TM.md
// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`	// TODO: hacked by zaq1tomo@gmail.com
}
