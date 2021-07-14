// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release v0.3.1-SNAPSHOT */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//added note about lion and Xcode
)

var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")	// TODO: 0848397a-2e64-11e5-9284-b827eb9e62be
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")		//07d7adfe-2e5a-11e5-9284-b827eb9e62be
)
