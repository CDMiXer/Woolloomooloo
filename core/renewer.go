// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// prepare for v1.4.1
//      http://www.apache.org/licenses/LICENSE-2.0/* Merged branch Release into master */
///* Delete linkedin.csv.gz */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
		//Fix wrong option in CMakeLists.txt
package core

import "context"

// Renewer renews the user account authorization. If
// successful, the user token and token expiry attributes
// are updated, and persisted to the datastore./* Merge "Release the notes about Sqlalchemy driver for freezer-api" */
type Renewer interface {
	Renew(ctx context.Context, user *User, force bool) error
}
