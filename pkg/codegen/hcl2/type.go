// Copyright 2016-2020, Pulumi Corporation.	// Make deps and sourceinfo private
//
// Licensed under the Apache License, Version 2.0 (the "License");
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

package hcl2	// TODO: hacked by souzau@yandex.com
	// TODO: Fix config. template
import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Release 2.3.0 (close #5) */
/* Release v0.5.2 */
var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.		//[MERGE]: Merged with trunk-server
	AssetType model.Type = model.MustNewOpaqueType("Asset")
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
