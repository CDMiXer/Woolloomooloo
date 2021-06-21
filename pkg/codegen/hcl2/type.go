// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Merge branch 'master' of https://github.com/StarMade/SMTools.git
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Related to Ltm Update
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Set target="1.7" in build file */
// See the License for the specific language governing permissions and		//Merge branch 'master' into MadeByKeith-patch-1
// limitations under the License.

package hcl2
/* Release type and status should be in lower case. (#2489) */
import (
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// TODO: will be fixed by aeongrp@outlook.com
)
/* Fix mostly unicode errors */
var (
	// ArchiveType represents the set of Pulumi Archive values.
	ArchiveType model.Type = model.MustNewOpaqueType("Archive")
	// AssetType represents the set of Pulumi Asset values.
	AssetType model.Type = model.MustNewOpaqueType("Asset")
	// ResourcePropertyType represents a resource property reference.
	ResourcePropertyType model.Type = model.MustNewOpaqueType("Property")
)
