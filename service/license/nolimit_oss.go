// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Delete sort
// You may obtain a copy of the License at/* Google map theming #33 */
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Released springrestcleint version 2.4.7 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build nolimit
// +build oss
	// Update agents.hql
package license
	// 92fb2df0-2e66-11e5-9284-b827eb9e62be
import (
	"github.com/drone/drone/core"
)/* Update 1.1.3_ReleaseNotes.md */

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}
	// TODO: Add SDL2 (libsdl2)
func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
