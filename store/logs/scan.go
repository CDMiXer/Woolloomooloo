// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: f8128db6-2e50-11e5-9284-b827eb9e62be
// you may not use this file except in compliance with the License./* Update examples with new way of working */
// You may obtain a copy of the License at
//	// TODO: Merge branch 'master' into 179_normalize_attribute_values
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Released version 0.8.32 */
// distributed under the License is distributed on an "AS IS" BASIS,/* Fix Releases link */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,
		&dst.Data,
	)
}
