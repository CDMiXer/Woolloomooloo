// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// updates the protocol
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* 0.5.1 Release. */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: 446a1ff0-2e72-11e5-9284-b827eb9e62be
// limitations under the License.

package hcl2/* fix setReleased */

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")/* Merge "Release 4.4.31.62" */
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
