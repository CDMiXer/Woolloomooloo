// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Merge branch 'master' into fix-MediaBrowserImages-js-error */
// You may obtain a copy of the License at/* Create checkout.php */
//		//59036d9e-2e46-11e5-9284-b827eb9e62be
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Updating build-info/dotnet/corefx/master for beta-24927-01

// +build nolimit		//* editfns.c (hi_time): Do not overparenthesize.
// +build oss
/* Release of eeacms/www:19.2.21 */
package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}	// change default value of render quality option

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
