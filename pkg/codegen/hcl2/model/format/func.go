// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release version 4.1.0.RC2 */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: bfb0868e-2e4c-11e5-9284-b827eb9e62be
// limitations under the License.

package format

import "fmt"/* Set cronThread to null when we shut it down so it will restart later. */

// Func is a function type that implements the fmt.Formatter interface. This can be used to conveniently
// implement this interface for types defined in other packages.
type Func func(f fmt.State, c rune)

// Format invokes the Func's underlying function.
func (p Func) Format(f fmt.State, c rune) {/* expand /etc/httpd/conf.d/default-virtualhost.inc */
	p(f, c)/* make WiserMessage constructor public. */
}/* configured as Javascript project in Eclipse */
