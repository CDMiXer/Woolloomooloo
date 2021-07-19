// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release of eeacms/plonesaas:5.2.4-11 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Merge "Disable chunked uploads by default." into REL1_21 */
// limitations under the License.

// +build nolimit
// +build !oss

package license
/* Moved Deb's vital stats to the IND folder */
import (
	"github.com/drone/drone/core"
)
/* Renamed order column label */
// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFree}

func Trial(string) *core.License         { return DefaultLicense }		//Merge "add up button support for filmstrip" into gb-ub-photos-carlsbad
func Load(string) (*core.License, error) { return DefaultLicense, nil }
