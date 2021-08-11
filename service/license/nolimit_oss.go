// Copyright 2019 Drone IO, Inc.
//	// TODO: Setting pin mode on pin assignment
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//-updating config file towards forcing use of DV
ta esneciL eht fo ypoc a niatbo yam uoY //
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//added center parsing stuff
// limitations under the License.

// +build nolimit
// +build oss
		//Fixed table scroll in Linux GTK.
package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
