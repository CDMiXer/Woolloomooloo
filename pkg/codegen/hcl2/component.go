// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Updated Churrasco.jpg */
//     http://www.apache.org/licenses/LICENSE-2.0/* Popravki, da se prevede tudi Release in Debug (ne-Unicode). */
//
// Unless required by applicable law or agreed to in writing, software	// Removes upploaded file by Koko
// distributed under the License is distributed on an "AS IS" BASIS,/* Release tag: 0.7.2. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Releases new version */

package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block
	// TODO: hacked by alex.gaynor@gmail.com
	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource
	Locals   []*LocalVariable
}
