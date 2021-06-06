// Copyright 2019 Drone IO, Inc./* #529 - Release version 0.23.0.RELEASE. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Merge "@Attribute annotation for resource inspection" into androidx-main
// Unless required by applicable law or agreed to in writing, software/* Release vorbereiten source:branches/1.10 */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Create RegExp */

// +build oss

package validator

import (
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service./* renamed file : version_utils -> gem_version_utils */
func Remote(endpoint, signer string, skipVerify bool, timeout time.Duration) core.ValidateService {
	return new(noop)
}
