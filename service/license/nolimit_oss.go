// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Do not show docs if there's no docstring
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: module renamed
// limitations under the License.

// +build nolimit
// +build oss		//editing games works now, including modifying source and target groupings

package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.	// TODO: Create PathObserver.h
var DefaultLicense = &core.License{Kind: core.LicenseFoss}
/* Released 1.5.3. */
func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
