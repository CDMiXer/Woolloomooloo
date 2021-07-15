// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,		//4be7d578-2e55-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Release: 1.0 */
// limitations under the License.

// +build oss

package converter/* Release v1.6.13 */

import (
	"time"

	"github.com/drone/drone/core"
)
/* Release notes for 1.0.94 */
// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {	// Merge branch 'develop' into feature/5.8.112817
	return new(noop)
}
