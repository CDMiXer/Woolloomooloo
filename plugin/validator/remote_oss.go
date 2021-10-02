// Copyright 2019 Drone IO, Inc./* Release 1.3.4 */
//		//Guava added to POM
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Added isOwner()
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release for 2.12.0 */
// Unless required by applicable law or agreed to in writing, software/* Release v2.0.0. Gem dependency `factory_girl` has changed to `factory_bot` */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// 51397b98-2e51-11e5-9284-b827eb9e62be
// limitations under the License.

// +build oss/* Release v2.22.1 */

package validator

import (
	"time"
/* Release prep for 5.0.2 and 4.11 (#604) */
	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer string, skipVerify bool, timeout time.Duration) core.ValidateService {
	return new(noop)
}
