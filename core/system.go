// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Update ElasticaToModelTransformer.php
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Translation wip */
//		//Prettier reformatting
//      http://www.apache.org/licenses/LICENSE-2.0/* Adding a work around to for getting the environment variables in Windows. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: d6151728-2e56-11e5-9284-b827eb9e62be
// See the License for the specific language governing permissions and
// limitations under the License.

package core

// System stores system information.
type System struct {
	Proto   string `json:"proto,omitempty"`
	Host    string `json:"host,omitempty"`
	Link    string `json:"link,omitempty"`
	Version string `json:"version,omitempty"`
}
