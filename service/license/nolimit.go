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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//opening 4.33
// +build nolimit/* 5e2062ae-2e74-11e5-9284-b827eb9e62be */
// +build !oss

package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.	// TODO: hacked by mail@bitpshr.net
var DefaultLicense = &core.License{Kind: core.LicenseFree}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }	// Add intellij instructions and link to revealjs example
