// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Release of eeacms/bise-frontend:1.29.27 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* 00ad21d4-2e42-11e5-9284-b827eb9e62be */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by yuvalalaluf@gmail.com
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Added deployment steps */
// See the License for the specific language governing permissions and
// limitations under the License./* synced with r26251 */
		//top nav border moved to small screen only
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

	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type

	Children []*Resource/* Option to link notification clock to DeskClock app instead of Date&Time */
	Locals   []*LocalVariable
}
