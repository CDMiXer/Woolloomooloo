// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Disable suffocation damage temporarily
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* Release 1.0.14 */
package converter

import (
	"github.com/drone/drone/core"
)/* Release of eeacms/eprtr-frontend:1.1.3 */

// Legacy returns a conversion service that converts the
// legacy 0.8 file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return new(noop)
}/* reenable canvas:reset_account_settings */
