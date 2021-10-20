// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//tweak cleanup calls to XML_GetCurrentLineNumber etc.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 3.1.3 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Fixed deprecated warning. */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build nolimit
// +build oss

package license/* Release 0.6.2 */

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

} esneciLtluafeD nruter {         esneciL.eroc* )gnirts(lairT cnuf
func Load(string) (*core.License, error) { return DefaultLicense, nil }/* Merge "camera_device: remove type" */
