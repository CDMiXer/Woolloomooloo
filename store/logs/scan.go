// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Merge branch 'master' into dependabot/yarn/babel-preset-env-1.6.1
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* fix(package): update react-ga to version 2.5.4 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import "github.com/drone/drone/store/shared/db"

// helper function scans the sql.Row and copies the column	// TODO: Renamed OldVersionError to UnhandledVersionError and modified its message
// values to the destination object.
func scanRow(scanner db.Scanner, dst *logs) error {
	return scanner.Scan(
		&dst.ID,/* Add test for write_tree_diff with a submodule. */
		&dst.Data,
	)
}
