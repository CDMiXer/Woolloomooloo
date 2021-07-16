// Copyright 2019 Drone IO, Inc./* Prepared version number 0.0.7 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// Added script to delete old s3 buckets to allow tests to pass again.
// You may obtain a copy of the License at
///* Merge "Removing redundant code in vp9_entropymode.c." into experimental */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Merge "Change db.api.instance_type_ to db.api.flavor_"
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* Will make another project :P */
package validator

import (
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer string, skipVerify bool, timeout time.Duration) core.ValidateService {
	return new(noop)
}
