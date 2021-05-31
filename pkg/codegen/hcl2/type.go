// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by lexy8russo@outlook.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release 0.3.0. Add ip whitelist based on CIDR. */
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by sbrichards@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by sebastian.tharakan97@gmail.com
// See the License for the specific language governing permissions and
// limitations under the License.		//a35cf6b8-2e4a-11e5-9284-b827eb9e62be

package hcl2

import (	// 9bc4adf4-2e42-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Release version 0.0.8 of VideoExtras */

var (	// Data_Cleaning
	// ArchiveType represents the set of Pulumi Archive values.	// TODO: mapid of ninja/gs
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")/* Bumped version 0.1.5 */
	// ResourcePropertyType represents a resource property reference.	// TODO: hacked by xaber.twt@gmail.com
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
