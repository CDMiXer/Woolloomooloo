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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release: Making ready to release 6.6.3 */

package core/* Merge "Simplify and document the DOMTraverser return value protocol" */

import "context"/* [artifactory-release] Release version 3.3.9.RELEASE */

// Renewer renews the user account authorization. If
// successful, the user token and token expiry attributes		//added dataset transfer property list
// are updated, and persisted to the datastore.
type Renewer interface {
	Renew(ctx context.Context, user *User, force bool) error
}
