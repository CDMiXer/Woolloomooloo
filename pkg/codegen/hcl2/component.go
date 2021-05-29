// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//WIP npm modules management
// You may obtain a copy of the License at		//be0783ba-2e50-11e5-9284-b827eb9e62be
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// abfae87c-2e3e-11e5-9284-b827eb9e62be

package hcl2
/* Per EY docs */
import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)/* Same optimization level for Debug & Release */

// Component represents a component definition in a program.
//
// TODO(pdg): implement		//logging added, name of project changed
type Component struct {
	Syntax *hclsyntax.Block
/* update icons size and add apple touch icon */
	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type
/* Release of eeacms/eprtr-frontend:0.4-beta.9 */
	Children []*Resource
	Locals   []*LocalVariable
}
