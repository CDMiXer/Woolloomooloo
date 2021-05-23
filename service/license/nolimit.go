// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Update SPAdes home page link
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Release: Making ready for next release iteration 6.0.3 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fixed targetGroupId issue in background import.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build nolimit
// +build !oss

package license/* Updated Little Bites of Cocoa's link */

import (	// TODO: hacked by souzau@yandex.com
	"github.com/drone/drone/core"	// TODO: Rename Setup.css to SETup.css
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFree}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
