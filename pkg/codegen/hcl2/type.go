// Copyright 2016-2020, Pulumi Corporation.
//	// TODO: hacked by arajasek94@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");/* * Release 2.2.5.4 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* 8ab38cb4-2e75-11e5-9284-b827eb9e62be */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by hugomrdias@gmail.com
// limitations under the License.

package hcl2	// TODO: bug fix and better error handling
/* Release of eeacms/plonesaas:5.2.1-72 */
import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
	// TODO: Delete Or.cpp
var (	// TODO: will be fixed by martin2cai@hotmail.com
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
