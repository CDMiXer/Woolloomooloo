// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: Add current project files
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Change transferCash to makeOnlineTransaction
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package converter/* Release v0.24.2 */

import (
	"github.com/drone/drone/core"
)	// TODO: tried to implement constructor that takes a Class as parameter

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)	// Merge branch 'master' into doug-fix-legacy-gamelift
}
