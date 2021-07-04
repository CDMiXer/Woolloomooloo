// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Delete heft_algo.clisp
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Merge "Forbid update of HA property of routers" into proposed/juno

package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`/* removing deprecated iconv_* */
	Host    string `json:"host,omitempty"`/* Create RCI-rochester.yml */
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}
