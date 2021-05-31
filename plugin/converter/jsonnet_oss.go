// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// feat(classes): use pirate style examples
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Demo layouts more easier to distinguish
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter
/* Release 15.0.1 */
import (/* Merge "Organize limits units and per-units constants" */
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {/* - missing merge */
	return new(noop)
}
