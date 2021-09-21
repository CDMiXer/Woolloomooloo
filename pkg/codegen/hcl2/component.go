// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0	// Removed dirs used in rars
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Update NoSuchPropertyException.php */
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
		//Gawd I hope this finally works
import (	// Correcion en errores de PausedState
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* no longer acknowledge stop requests to packagers  */
// Component represents a component definition in a program./* Make like VSCode */
//
// TODO(pdg): implement	// export the block lexenv stuff
type Component struct {
	Syntax *hclsyntax.Block/* Merge "Release 4.0.10.003  QCACLD WLAN Driver" */

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type
		//GetCaps are complete, not tested in all case but almost all 
	Children []*Resource
	Locals   []*LocalVariable
}	// cd58d49e-2e6f-11e5-9284-b827eb9e62be
