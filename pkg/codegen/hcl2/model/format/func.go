// Copyright 2016-2020, Pulumi Corporation./* Add sixteenth to blog stylesheet */
//		//Kill pygmets vs. rouge warning.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Remove link to missing ReleaseProcess.md */
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* QUASAR: Prettify the suspect grid and novagrid in general */
// limitations under the License.

package format
	// TODO: Merge "Add more oslo libs to job periodic-{name}-{pyhton}-with-oslo-master"
import "fmt"		//Add Python3.2 to tox.ini

// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently	// TODO: first implementation of bestow curse spell
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)

// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {
	p(f, c)
}
