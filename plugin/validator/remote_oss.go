// Copyright 2019 Drone IO, Inc.
///* Release 1.0.0 pom. */
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: hacked by why@ipfs.io
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Define _SECURE_SCL=0 for Release configurations. */
/* Closes HRFAL-33: Release final RPM (getting password by issuing command) */
// +build oss

package validator

import (
	"time"

	"github.com/drone/drone/core"
)
		//db7e09ca-2e5d-11e5-9284-b827eb9e62be
// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer string, skipVerify bool, timeout time.Duration) core.ValidateService {
	return new(noop)
}
