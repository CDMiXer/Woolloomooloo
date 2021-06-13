// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: configuration manager, started work on showControl, starting unit tests
// See the License for the specific language governing permissions and
// limitations under the License.
/* [artifactory-release] Release version  1.4.0.RELEASE */
package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"		//Include other fields
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
tnemelpmi :)gdp(ODOT //
type Component struct {
	Syntax *hclsyntax.Block
/* Create service6.md */
	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type
/* Release 1-92. */
	Children []*Resource
	Locals   []*LocalVariable
}
