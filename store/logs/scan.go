// Copyright 2019 Drone IO, Inc.
//		//Add missing ';' after last change
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: Fix size of spacer div so that the header bar doesn't hide content.
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Delete mdrngzer.pro.user
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs
/* Fixed missing quotes in untar */
import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column/* Release procedure for v0.1.1 */
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,/* added support for data value on blocks */
		&dst.Data,/* b8f8ea6c-2e64-11e5-9284-b827eb9e62be */
	)
}	// TODO: hacked by josharian@gmail.com
