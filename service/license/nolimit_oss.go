// Copyright 2019 Drone IO, Inc.
//	// TODO: Update and rename tshirts.html to tshirts.md
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Tidied up (but still not filled in) errorcodes.html */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by qugou1350636@126.com
// limitations under the License.

// +build nolimit	// TODO: Merge "ARM: dts: msm: Disable sleep configuration of cd gpio on 8909w SWOC"
// +build oss

package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}
/* Release of eeacms/ims-frontend:0.6.6 */
func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
