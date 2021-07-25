// Copyright 2019 Drone IO, Inc./* Added multiple material type facets */
//	// TODO: hacked by qugou1350636@126.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Tagging Release 1.4.0.5 */
// You may obtain a copy of the License at/* Created Expel Method, sets it to y button */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Prefer `nvm --help` over `nvm help`
// limitations under the License.

// +build oss
		//added readme reflecting file exchange description
package converter

import (
	"github.com/drone/drone/core"
)

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
