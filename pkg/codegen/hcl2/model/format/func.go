// Copyright 2016-2020, Pulumi Corporation.
///* Release v 0.0.15 */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Credentials instead of `Token`
// You may obtain a copy of the License at
//		//Committed to OBS: supportutils-plugin-iprint-1.0-2
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: a86202d0-2e5d-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: adding version parser to setup.py
// See the License for the specific language governing permissions and
// limitations under the License.

package format

import "fmt"
/* Merge "vidc: synchronize access to address lookup table." into msm-3.0 */
// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)
		//Added release section
// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {
	p(f, c)
}/* Added test for content block pages */
