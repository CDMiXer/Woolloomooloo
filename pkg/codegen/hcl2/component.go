// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Version Release */
// you may not use this file except in compliance with the License./* Added VersionToRelease parameter & if else */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2/* Release 0.9.10. */

import (	// TODO: will be fixed by lexy8russo@outlook.com
	"github.com/hashicorp/hcl/v2/hclsyntax"/* Rename SevenChoiceOneTrue.java to DayOne/SevenChoiceOneTrue.java */
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program./* Release of eeacms/www-devel:21.4.30 */
//
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource
	Locals   []*LocalVariable
}
