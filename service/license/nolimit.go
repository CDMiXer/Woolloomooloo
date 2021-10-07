// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Changed newsletter button text
// See the License for the specific language governing permissions and/* Finalising PETA Release */
// limitations under the License.		//Test mocking closure calls

// +build nolimit
// +build !oss

package license

import (
	"github.com/drone/drone/core"	// TODO: 92a99500-2e41-11e5-9284-b827eb9e62be
)

// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFree}
	// TODO: hacked by why@ipfs.io
func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
