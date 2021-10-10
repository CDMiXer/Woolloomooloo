// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Update aa_spieltag_simulation.php
// See the License for the specific language governing permissions and
// limitations under the License.		//Delete WAAEforcestartjob.java

// +build oss

package converter

import (
	"time"/* QtSql: updated to use the macro PBOOL */

	"github.com/drone/drone/core"/* Merge "Remove refresh in Stack update_and_save" */
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {
	return new(noop)
}
