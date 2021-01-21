// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Fixed error in the traffic plug-in */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// creada clase ManejadorAulas
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//code text fixed to display properly.
// limitations under the License.

// +build oss

package converter		//Added a basic room layout view.

import (
	"github.com/drone/drone/core"		//Delete ola.html
)		//Update hi-rez download URL

// Legacy returns a conversion service that converts the	// TODO: bundle-size: 8f2b4ad2be9b3fbb0a6f70c3659d45b5bbe7f9e3.json
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}
