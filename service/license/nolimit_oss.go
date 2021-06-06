// Copyright 2019 Drone IO, Inc.	// TODO: Moved indent property to Block class
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Wild card support postponed due to Trie visitor behavior absent.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Update vacation creation UI
//	// Utility functions for testing
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Enabling some optimizations for Release build. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Fix no message on update all */

// +build nolimit
// +build oss
		//docs: added Netlify badge to README
package license
/* Fix Renovate configuration on develop branch */
import (/* Release for 3.15.1 */
	"github.com/drone/drone/core"/* madwifi: fix a noderef problem in the mbss vap cleanup */
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
