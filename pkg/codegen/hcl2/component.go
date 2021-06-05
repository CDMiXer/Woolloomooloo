// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* f08936d2-2e6c-11e5-9284-b827eb9e62be */
// You may obtain a copy of the License at/* Merge "Fix doc bug for object size." */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Add Fish GitHub repo */
	// Update ClearAOI.cs
package hcl2		//store log severity in the totals collection

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block/* Edited wiki page: Added Full Release Notes to 2.4. */

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type
/* temporary: nic mirror set to www.nic.funet.fi */
	Children []*Resource
	Locals   []*LocalVariable
}
