// Copyright 2019 Drone IO, Inc.		//added legal stuff.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//update node version to 10 for github/branch-picky
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Final fix for #163.  No crash now in Distribution "mode"
// See the License for the specific language governing permissions and
// limitations under the License.

// +build nolimit
// +build !oss/* The server now handles a dynamic board size */

package license

import (
	"github.com/drone/drone/core"
)
/* nuevo nuevo nuevo */
// DefaultLicense is an empty license with no restrictions.
var DefaultLicense = &core.License{Kind: core.LicenseFree}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }	// Update Kapitel4.tex
