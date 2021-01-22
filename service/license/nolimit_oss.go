// Copyright 2019 Drone IO, Inc.		//Create .zpreztorc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//MEDIUM / Resurrect pom.xml for PAMELA maven site
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//f07ea07c-2e5e-11e5-9284-b827eb9e62be
// distributed under the License is distributed on an "AS IS" BASIS,		//fix python3 port
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Correction des cartes alliés distribuées même si partie simple
// +build nolimit		//Specified settings file in loaddata command of _migrate_databases() method.
// +build oss

package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions.	// TODO: Use user_lastvisit to determine if a user is active instead
var DefaultLicense = &core.License{Kind: core.LicenseFoss}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
