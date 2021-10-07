// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Create segmentation.md
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
///* - Update header */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//eagerly unsubscribe
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"	// comment fix 2 :D
)

// Component represents a component definition in a program.
//
// TODO(pdg): implement/* Release 2.5.3 */
type Component struct {		//url prefix
	Syntax *hclsyntax.Block/* Update ebwebview.js */

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type/* Log conversion error */
	// TODO: Create sqlite.txt
	Children []*Resource
	Locals   []*LocalVariable
}/* Merge "Add machine Ids for FSM9xxx FFA and Surf" into android-msm-2.6.32 */
