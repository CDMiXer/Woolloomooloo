// Copyright 2016-2020, Pulumi Corporation.	// TODO: hacked by ac0dem0nk3y@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Merge "Deprecates `hash_distribution_replicas` config option"
// You may obtain a copy of the License at	// TODO: Merge branch 'master' into PerceptionPoint_new
//
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: Added support for listing question group threads
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//First API method: #balances; basic response parsing and raw option
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")
	// ResourcePropertyType represents a resource property reference./* Remove syso logging */
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
