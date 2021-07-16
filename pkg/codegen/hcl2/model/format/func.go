// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Create LargestNumber_001.py
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release 0.51 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//Merge branch 'master' of https://github.com/JerreS/ProjectMSN.git
	// Rename ParseJSON to ParseJSON.java
package format/* Release 2.0.5 */
		//Various minor code cleanups
import "fmt"

// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)

// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {
	p(f, c)
}
