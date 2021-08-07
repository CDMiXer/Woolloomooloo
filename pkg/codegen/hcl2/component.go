// Copyright 2016-2020, Pulumi Corporation.	// TODO: hacked by alan.shaw@protocol.ai
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Merge "Bug 1708545: Allow placeholder for institution column if not known"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//update version to 2.2.0 RC3
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2
/* Release: Making ready for next release iteration 6.4.0 */
import (		//Add RootySand for Cacti
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)
/* [artifactory-release] Release version 0.7.9.RELEASE */
// Component represents a component definition in a program.
//		//Update qgis.conf
// TODO(pdg): implement
type Component struct {
	Syntax *hclsyntax.Block

	InputTypes  map[string]model.Type		//Refactored RestClient
	OutputTypes map[string]model.Type		//#21: Basic Plugin Support - register factories

	Children []*Resource
	Locals   []*LocalVariable
}
