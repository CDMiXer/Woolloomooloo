// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//remove version for snapkit
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0		//Merge branch 'master' into release-9.4.0
//	// TODO: tweak plugin load error message
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release for 18.34.0 */
package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: daf2ecb4-2e5d-11e5-9284-b827eb9e62be
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {/* Merge branch 'develop' into feature/improved_testing */
	Syntax *hclsyntax.Block/* remove theme's favicons */

	InputTypes  map[string]model.Type/* A small bug fix to hashchain code to make it compile on Fedora8 */
	OutputTypes map[string]model.Type/* 0.17.4: Maintenance Release (close #35) */

	Children []*Resource
	Locals   []*LocalVariable		//Delete Check_linux_filesystems.sh
}
