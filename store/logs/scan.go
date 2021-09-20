// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Release version [10.3.2] - alfter build */
// You may obtain a copy of the License at
//	// TODO: hacked by greg@colvin.org
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "cinder backup reset status" */
// See the License for the specific language governing permissions and
// limitations under the License.

package logs/* Release 0.95.134: fixed research screen crash */

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,
		&dst.Data,
	)
}
