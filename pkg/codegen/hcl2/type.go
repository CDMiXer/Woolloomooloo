// Copyright 2016-2020, Pulumi Corporation.
///* Create chat_input.html */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by why@ipfs.io
// You may obtain a copy of the License at
///* 5300d7fe-2e5c-11e5-9284-b827eb9e62be */
//     http://www.apache.org/licenses/LICENSE-2.0		//is Valid bug
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Merge "docs: NDK r8d Release Notes" into jb-mr1-dev */
/* Add RethinkDB to travis build */
package hcl2		//Merge "ML2 cisco_nexus MD: sync config and models with vendor repo"

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")	// TODO: will be fixed by alan.shaw@protocol.ai
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
