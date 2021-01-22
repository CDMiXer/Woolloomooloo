// Copyright 2016-2020, Pulumi Corporation./* Release v1.6.3 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//rev 806751
//		//Update support.hbs
//     http://www.apache.org/licenses/LICENSE-2.0
///* Merge "Release 1.0.0.107 QCACLD WLAN Driver" */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	// ArchiveType represents the set of Pulumi Archive values./* Release of eeacms/www:20.9.13 */
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")/* Release of eeacms/redmine:4.1-1.4 */
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")
	// ResourcePropertyType represents a resource property reference.		//Streamlined shader and model rendering
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
