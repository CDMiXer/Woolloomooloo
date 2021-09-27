// Copyright 2019 Drone IO, Inc.	// Merge "Update basic listen."
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* add link to tracker */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release for 2.5.0 */
// See the License for the specific language governing permissions and	// TODO: Update adsense.php
// limitations under the License.
/* Release Neo4j 3.4.1 */
// +build oss
/* Updated so the static files come from one site. */
package admission
/* Release areca-5.3.5 */
import "github.com/drone/drone/core"

// External is a no-op admission controller
func External(string, string, bool) core.AdmissionService {
	return new(noop)
}
