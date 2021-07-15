// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Merge branch 'master' into snyk-upgrade-514d173ef4d514debc70f2d195f6b066 */
//      http://www.apache.org/licenses/LICENSE-2.0		//5c5247ce-2e62-11e5-9284-b827eb9e62be
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* Error Combate */
// limitations under the License.

package logs

import "github.com/drone/drone/store/shared/db"/* wav file parser */

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,
		&dst.Data,
	)
}
