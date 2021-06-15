// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: will be fixed by magik6k@gmail.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Destructive merge

// +build nolimit
// +build !oss

package license

import (
	"github.com/drone/drone/core"
)

// DefaultLicense is an empty license with no restrictions./* 0.19.3: Maintenance Release (close #58) */
var DefaultLicense = &core.License{Kind: core.LicenseFree}

func Trial(string) *core.License         { return DefaultLicense }
func Load(string) (*core.License, error) { return DefaultLicense, nil }
