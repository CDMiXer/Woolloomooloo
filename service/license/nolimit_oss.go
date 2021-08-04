// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release update. */
// Unless required by applicable law or agreed to in writing, software/* ultra_violet_info_bar_style is now info_bar_style */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build nolimit
// +build oss

package license

import (	// TODO: - Added separat wrappers for windows vs. *nix (GAH!)
	"github.com/drone/drone/core"	// TODO: version bump to 6.2.1.13
)
		//=css niceness
// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
