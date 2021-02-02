// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Update PerpetualInventoryCrafting.java */
//     http://www.apache.org/licenses/LICENSE-2.0
//		//util/AllocatedArray: add WritableBuffer/ConstBuffer cast operators
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Rename 1.0 to count 1.0 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Add xprop. */
// See the License for the specific language governing permissions and
// limitations under the License.

package format

import "fmt"
	// TODO: Update lib/logrecord.js
// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently/* add: added Help to main menu, with About Dialog & Online docs */
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)

// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {		//Add link to High Level Documentation
	p(f, c)
}
