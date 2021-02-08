// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: will be fixed by why@ipfs.io
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// Help Command is polished
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* Release tag: 0.7.2. */
// +build oss

package validator

import (
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer string, skipVerify bool, timeout time.Duration) core.ValidateService {
	return new(noop)/* Release areca-5.5.6 */
}
