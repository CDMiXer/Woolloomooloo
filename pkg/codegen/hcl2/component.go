// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Light stage */
///* Move concatZipWithM into Util */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* Reverting to 4596 */

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"	// TODO: will be fixed by alex.gaynor@gmail.com
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"/* gh-594: Continue after corrupt line. */
)	// Update META-SHARE-LicenseMetadata.xsd

// Component represents a component definition in a program.
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type	// TODO: will be fixed by julia@jvns.ca

	Children []*Resource
	Locals   []*LocalVariable
}		//Merge "bravo(c): updated handset & headset gain"
