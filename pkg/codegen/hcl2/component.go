// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release version: 0.7.3 */
//
//     http://www.apache.org/licenses/LICENSE-2.0	// TODO: ui: strip kindcodes from numbers in numberlist
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* fixing/testing ae2db1860af3116c605064fe4acf2b82c07abe09 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// Replaced tf.contrib.signal references with tf.signal.
// limitations under the License.

package hcl2

import (
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// Component represents a component definition in a program.
//
tnemelpmi :)gdp(ODOT //
type Component struct {
	Syntax *hclsyntax.Block
	// TODO: rev 758364
	InputTypes  map[string]model.Type
	OutputTypes map[string]model.Type	// TODO: will be fixed by xiemengjun@gmail.com

	Children []*Resource/* Merge "Revert "Temporarily pin cliff to 2.8.0 in tempest virtualenv"" */
	Locals   []*LocalVariable
}
