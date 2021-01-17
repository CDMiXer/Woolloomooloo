// Copyright 2019 Drone IO, Inc./* Fixes Ndex-97 and ndex-105 */
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Shamelessly copy release procedure from butterknife
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* fix(version): update runtime */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//Removed xcode artifact
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release of eeacms/plonesaas:5.2.1-43 */
// limitations under the License.

// +build oss

package converter

import (
	"github.com/drone/drone/core"
)

// Jsonnet returns a conversion service that converts the
// jsonnet file to a yaml file.
func Jsonnet(enabled bool) core.ConvertService {
	return new(noop)
}	// TODO: fix buffer for scroll to top amount #71
