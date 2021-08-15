// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Next Release Version Update */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Sort comments descending when shown
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* [Issue-100] To Support carbon compatible hive syntax for carbon tables (#423) */
// See the License for the specific language governing permissions and
// limitations under the License.

package format
/* Update compatibility info */
import "fmt"

// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)		//Update Elbow Code

// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {
	p(f, c)	// TODO: Update lawyer mailer spec to skip assertion
}
