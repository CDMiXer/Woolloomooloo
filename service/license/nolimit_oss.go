// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release version 3.1.0.M1 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* reverse to old code */

// +build nolimit
// +build oss

package license	// TODO: hacked by timnugent@gmail.com
/* Fixes Issue #33 */
import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
